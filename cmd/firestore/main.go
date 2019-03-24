package main

import (
	"context"
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

			fn := c.CreateSnapshotsWriter(*river, ch)
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

func (c *cache) CreateSnapshotsWriter(river River, ch chan *gauge.Snapshot) func(ctx context.Context, d *daemon.Supervisor) error {
	return func(ctx context.Context, d *daemon.Supervisor) error {
		for {
			var snap *gauge.Snapshot
			select {
			case <-ctx.Done():
				return nil
			case snap = <-ch:
			}

			span := report.StartSpan("snapshot.applied",
				report.TraceID(snap.CorrelationID), report.ParentSpanID(snap.CausationID))
			span = span.Field("section_uuid", river.Section.UUID)

			// TODO: process snapshot

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
	urls := []string{
		s.Station.DataURL,
		s.Station.AliasURL,
		s.Station.HumanURL,
	}
	for _, url := range urls {
		chs, ok := c.SnapRoute[url]
		if ok {
			for _, ch := range chs {
				ch <- s
			}
		}
	}

	return nil
}
