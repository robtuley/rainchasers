package main

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/internal/queue"
	"github.com/rainchasers/report"
)

func main() {
	d := daemon.New("firestore")
	url := "https://app.rainchasers.com/catalogue.json"
	app := &cache{
		ProjectID: os.Getenv("PROJECT_ID"),
		TopicName: os.Getenv("PUBSUB_TOPIC"),
		Rivers:    NewRiverCache(url),
		Log:       d.Logger,
	}

	d.Run(context.Background(), app.PollForRiverCatalogueChanges)
	d.Run(context.Background(), app.SubscribeToSnapshotUpdates)
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
	Rivers    *RiverCache
	Log       *report.Logger
}

func (c *cache) PollForRiverCatalogueChanges(ctx context.Context, d *daemon.Supervisor) error {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		isChanged, span := c.Rivers.Update(ctx)
		if isChanged && span.Err() == nil {
			span = span.FollowedBy(c.OnRiverCatalogueChange(ctx))
		}
		d.Trace(span)
		if err := span.Err(); err != nil {
			return err
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *cache) OnRiverCatalogueChange(ctx context.Context) report.Span {
	span := report.StartSpan("firestore.rivers.changed")
	if c.ProjectID == "" {
		return span.End()
	}

	// connect to firestore
	client, err := firestore.NewClient(ctx, c.ProjectID)
	if err != nil {
		return span.End(err)
	}
	defer client.Close()

	// batch together writing all rivers
	collection := client.Collection("rivers")
	batch := client.Batch()
	c.Rivers.Each(func(s Section) bool {
		batch = batch.Set(collection.Doc(s.UUID), s)
		return true
	})

	// write batch to firestore
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	_, err = batch.Commit(ctx)
	if err != nil {
		return span.End(err)
	}

	return span.End()
}

func (c *cache) SubscribeToSnapshotUpdates(ctx context.Context, d *daemon.Supervisor) error {
	topic, span := queue.New(ctx, c.ProjectID, c.TopicName)
	d.Trace(span)
	if err := span.Err(); err != nil {
		return err
	}
	defer topic.Stop()

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
