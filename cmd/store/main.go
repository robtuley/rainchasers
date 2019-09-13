package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/rainchasers/content"
	"github.com/rainchasers/content/internal/daemon"
	"github.com/rainchasers/content/internal/gauge"
	"github.com/rainchasers/content/internal/queue"
	"github.com/rainchasers/content/internal/river"
	"github.com/rainchasers/report"
)

func main() {
	d := daemon.New("firestore")
	app := &cache{
		ProjectID:     os.Getenv("PROJECT_ID"),
		TopicName:     os.Getenv("PUBSUB_TOPIC"),
		AlgoliaAppID:  os.Getenv("ALGOLIA_APP_ID"),
		AlgoliaAPIKey: os.Getenv("ALGOLIA_API_KEY"),
		ReadyC:        make(chan struct{}),
		Log:           d.Logger,
		SnapRoute:     make(map[string][]chan *gauge.Snapshot),
	}

	d.Run(context.Background(), app.Init)
	d.Run(context.Background(), app.SubscribeToSnapshots)
	d.CloseAfter(24 * time.Hour)

	d.Wait()
	if err := d.Err(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

type cache struct {
	ProjectID     string
	TopicName     string
	AlgoliaAppID  string
	AlgoliaAPIKey string
	ReadyC        chan struct{}
	Log           *report.Logger
	Writer        *FireWriter
	SnapRoute     map[string][]chan *gauge.Snapshot
}

func (c *cache) Init(ctx context.Context, d *daemon.Supervisor) error {
	// quit before any firestore prep if in dry run
	if c.ProjectID == "" {
		close(c.ReadyC)
		return nil
	}

	// connect to firestore
	fw, span := NewFireWriter(c.ProjectID, c.AlgoliaAppID, c.AlgoliaAPIKey)
	d.Trace(span)
	if err := span.Err(); err != nil {
		return err
	}
	c.Writer = fw

	// update catalogue in firestore (rate limited)
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
updateLoop:
	for _, s := range content.Sections {
		// get firestore info for the section
		// (& update if necessary)
		record, span := c.Writer.LoadAndUpdate(ctx, s)
		d.Trace(span)
		if err := span.Err(); err != nil {
			return err
		}

		// if calibration exists then launch goroutine
		// to listen to snapshots and update river
		calibrations, isCalibrated := content.Calibrations[s.UUID]
		if isCalibrated {
			ch := make(chan *gauge.Snapshot)

			// add to routing table
			for _, m := range calibrations {
				c.SnapRoute[m.URL] = append(c.SnapRoute[m.URL], ch)
			}

			fn := c.CreateSnapshotsWriter(*record, calibrations, ch)
			d.Run(context.Background(), fn)
		}

		select {
		case <-ctx.Done():
			break updateLoop
		case <-ticker.C:
		}
	}

	close(c.ReadyC)
	return nil
}

func (c *cache) CreateSnapshotsWriter(record Record, calibrations []river.Calibration, ch chan *gauge.Snapshot) func(ctx context.Context, d *daemon.Supervisor) error {
	return func(ctx context.Context, d *daemon.Supervisor) error {
		// the calibrations may have changed on previously inited measures
		// that have been pulled from firestore, so reset the calibrations
		// stored against each measure. if a calibration no longer exists, the
		// measure should be deleted.
		for i := range record.Measures {
			cal, exists := findCalibrationForStation(calibrations, record.Measures[i].Station)
			if exists {
				record.Measures[i].Calibration = cal
			} else {
				// delete this index from the slice
				record.Measures = append(record.Measures[:i], record.Measures[i+1:]...)
			}
		}

		// now we have updated the measures, we start building the map of where to
		// route snapshots too. Where a snapshot misses (i.e. a new measure), this
		// is lazy-inited as we need to wait for a snapshot with the station info in
		aliasURLToIndex := make(map[string]int)
		for i, m := range record.Measures {
			aliasURLToIndex[m.Station.AliasURL] = i
		}

	nextSnapshot:
		for {
			var snap *gauge.Snapshot
			ticker := time.NewTicker(4 * time.Hour)
			select {
			case <-ctx.Done():
				return nil
			case <-ticker.C:
				// if no snapshot received for some time there is
				// some sort of upstream problem
				c.Log.Action("snapshot.missing", report.Data{
					"section_uuid": record.Section.UUID,
				})
				ticker.Stop()
				continue nextSnapshot
			case snap = <-ch:
			}
			ticker.Stop()

			// route snap to existing or create new measure
			index, ok := aliasURLToIndex[snap.Station.AliasURL]
			if !ok {
				// this must be the first snapshot of a lazy-inited measure
				// so we need to search for the calibration and then setup
				// the new measure ready to receive further snapshots
				cal, exists := findCalibrationForStation(calibrations, snap.Station)
				if !exists {
					// this must be code logic as this routine should only receive
					// snaps that have a calibration for (even if that calibration is empty)
					msg := record.Section.UUID + " with snap " + snap.Station.AliasURL
					return errors.New("incorrectly routed snapshot: " + msg)
				}

				// append a new measure to the river with calibration and station
				// (readings will be added in the normal snapshot processing later)
				index = len(record.Measures)
				record.Measures = append(record.Measures, Measure{
					Station:     snap.Station,
					Calibration: cal,
					// no readings
				})
				aliasURLToIndex[snap.Station.AliasURL] = index
			}

			// now we know the index we're putting this snapshot into (and have
			// created a placeholder for it if it didn't exist), we merge in the
			// snapshot readings with the existing measure ones and save if changed
			span := report.StartSpan("snapshot.saved",
				report.TraceID(snap.CorrelationID), report.ParentSpanID(snap.CausationID))
			span = span.Field("section_uuid", record.Section.UUID)
			span = span.Field("alias_url", snap.Station.AliasURL)

			m := record.Measures[index]
			// checksum only the readings and station (as they are being potentially
			// changed). If checksum the whole measure, the processedTim field will
			// force a push on every snapshot received.
			prevChecksum := checksum(m.Readings, m.Station)

			// merge snapshot into the measure
			m.Readings = merge(m.Readings, snap.Readings)
			expiry := time.Now().Add(-4 * 24 * time.Hour)
			removeOlderThan(expiry, &m.Readings)
			m.Station = snap.Station
			if prevChecksum == checksum(m.Readings, m.Station) {
				// measure has not changed wait for next one
				continue nextSnapshot
			}
			m.ProcessedTime = snap.ProcessedTime
			record.Measures[index] = m

			// use updated measure to re-calulate level state
			record.Level = m.LatestLevel()

			// write the update to firestore
			wSpan := c.Writer.Store(ctx, &record)
			span = span.Child(wSpan)
			c.Log.Trace(span.End())
		}
	}
}

func (c *cache) SubscribeToSnapshots(ctx context.Context, d *daemon.Supervisor) error {
	// wait for init
	select {
	case <-ctx.Done():
		return nil
	case <-c.ReadyC:
	}

	// connect to pubsub
	topic, span := queue.New(ctx, c.ProjectID, c.TopicName)
	d.Trace(span)
	if err := span.Err(); err != nil {
		return err
	}
	defer topic.Stop()

	// subscribe!
	return topic.Subscribe(ctx, "", c.SnapshotRouter)
}

// only return error if want message redelivered, otherwise deal with it locally
func (c *cache) SnapshotRouter(ctx context.Context, err error, s *gauge.Snapshot) error {
	if err != nil {
		c.Log.Action("snapshot.corrupted", report.Data{
			"error": err.Error(),
		})
		return nil // error with decoding so do not retry delivery
	}

	// search any of data URL, alias URL, or human URL to route to
	// (using a map to remove dups between urls types)
	urls := make(map[string]bool)
	urls[s.Station.DataURL] = true
	urls[s.Station.AliasURL] = true
	urls[s.Station.HumanURL] = true
	for url := range urls {
		chs, ok := c.SnapRoute[url]
		if ok {
			for _, ch := range chs {
				ch <- s
			}
		}
	}

	return nil
}

func findCalibrationForStation(calibrations []river.Calibration, station gauge.Station) (cal river.Calibration, exists bool) {
	for _, c := range calibrations {
		if c.URL == station.DataURL {
			exists = true
			cal = c
		}
		if c.URL == station.AliasURL {
			exists = true
			cal = c
		}
		if c.URL == station.HumanURL {
			exists = true
			cal = c
		}
	}
	return
}
