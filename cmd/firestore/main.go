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
		// update section info in firestore and launcg goroutine
		// to listen for snaps
		_, span := c.Writer.LoadAndUpdate(ctx, s)
		d.Trace(span)
		if err := span.Err(); err != nil {
			return err
		}

		// TODO launch goroutine to watch any section with
		// a calibration

		select {
		case <-ctx.Done():
			break updateLoop
		case <-ticker.C:
		}
	}

	close(c.ReadyC)
	return nil
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
	return topic.Subscribe(ctx, "", c.OnReceivedSnapshot)
}

// only return error if want message redelivered, otherwise deal with it locally
func (c *cache) OnReceivedSnapshot(ctx context.Context, err error, s *gauge.Snapshot) error {
	if err != nil {
		c.Log.Action("snapshot.corrupted", report.Data{
			"error": err.Error(),
		})
		return nil // error with decoding so do not retry delivery
	}

	span := report.StartSpan("snapshot.processed",
		report.TraceID(s.CorrelationID), report.ParentSpanID(s.CausationID))

	// TODO: process snaps

	c.Log.Trace(span.End())
	return nil
}
