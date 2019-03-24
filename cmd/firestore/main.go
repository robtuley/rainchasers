package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/internal/queue"
	"github.com/rainchasers/report"
)

func main() {
	d := daemon.New("firestore")
	app := &cache{
		ProjectID: os.Getenv("PROJECT_ID"),
		TopicName: os.Getenv("PUBSUB_TOPIC"),
		ReadyC:    make(chan struct{}),
		Log:       d.Logger,
		SnapRoute: make(map[string][]chan *gauge.Snapshot),
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
	ProjectID string
	TopicName string
	ReadyC    chan struct{}
	Log       *report.Logger
	Writer    *FireWriter
	SnapRoute map[string][]chan *gauge.Snapshot
}

func (c *cache) Init(ctx context.Context, d *daemon.Supervisor) error {
	// parse catalogue
	sections, span := parseCatalogue(ctx)
	d.Trace(span)
	if err := span.Err(); err != nil {
		return err
	}

	// quit before any firestore prep if in dry run
	if c.ProjectID == "" {
		close(c.ReadyC)
		return nil
	}

	// connect to firestore
	fw, span := NewFireWriter(c.ProjectID)
	d.Trace(span)
	if err := span.Err(); err != nil {
		return err
	}
	c.Writer = fw

	// update catalogue in firestore (rate limited)
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
updateLoop:
	for _, s := range sections {
		// get firestore info for the section
		// (& update if necessary)
		river, span := c.Writer.LoadAndUpdate(ctx, s)
		d.Trace(span)
		if err := span.Err(); err != nil {
			return err
		}

		// if calibration exists then launch goroutine
		// to listen to snapshots and update river
		if len(s.Measures) > 0 {
			ch := make(chan *gauge.Snapshot)

			// add to routing table
			for _, m := range s.Measures {
				c.SnapRoute[m.URL] = append(c.SnapRoute[m.URL], ch)
			}

			fn := c.CreateSnapshotsWriter(*river, s.Measures, ch)
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

func (c *cache) CreateSnapshotsWriter(river River, calibrations []Calibration, ch chan *gauge.Snapshot) func(ctx context.Context, d *daemon.Supervisor) error {
	return func(ctx context.Context, d *daemon.Supervisor) error {
		aliasURLToIndex := make(map[string]int)
		for i, m := range river.Measures {
			aliasURLToIndex[m.Station.AliasURL] = i
		}

		for {
			var snap *gauge.Snapshot
			select {
			case <-ctx.Done():
				return nil
			case snap = <-ch:
			}

			// which measure index does this snapshot relate to, or is this
			// the very first snapshot?
			index, ok := aliasURLToIndex[snap.Station.AliasURL]
			if !ok {
				// this must be the first snapshot, search for an appropriate
				// calibration to map to it
				var cal Calibration
				for _, c := range calibrations {
					if c.URL == snap.Station.DataURL {
						cal = c
					}
					if c.URL == snap.Station.AliasURL {
						cal = c
					}
					if c.URL == snap.Station.HumanURL {
						cal = c
					}
				}
				if cal.URL == "" {
					msg := river.Section.UUID + " with snap " + snap.Station.AliasURL
					return errors.New("incorrectly routed snapshot: " + msg)
				}

				// now append the measure to the existing river with calibration and station
				// but no readings and log index for that
				index = len(river.Measures)
				river.Measures = append(river.Measures, Measure{
					Station:     snap.Station,
					Calibration: cal,
					// no readings
				})
				aliasURLToIndex[snap.Station.AliasURL] = index
			}

			// now we know the index we're putting this snapshot into (and have
			// created a placeholder for it if it didn't exist, we merge in the
			// snapshot readings
			span := report.StartSpan("snapshot.applied",
				report.TraceID(snap.CorrelationID), report.ParentSpanID(snap.CausationID))
			span = span.Field("section_uuid", river.Section.UUID)
			span = span.Field("alias_url", snap.Station.AliasURL)

			// TODO: merge in properly trimming on age, and only write if changed
			river.Measures[index].Readings = snap.Readings
			river.Measures[index].ProcessedTime = snap.ProcessedTime
			wSpan := c.Writer.Store(ctx, &river)
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
