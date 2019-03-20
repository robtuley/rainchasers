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
	url := "https://app.rainchasers.com/catalogue.json"
	app := &cache{
		ProjectID: os.Getenv("PROJECT_ID"),
		TopicName: os.Getenv("PUBSUB_TOPIC"),
		Rivers:    NewRiverCache(url),
	}

	d := daemon.New("firestore")
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
	span := report.StartSpan("rivers.firestore")
	// TODO: update firebase rivers
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

func (c *cache) OnReceivedSnapshot(ctx context.Context, err error, s *gauge.Snapshot) error {
	// TODO: process snaps, log error, etc
	// only return error if want message redelivered, otherwise deal with it locally
	return nil
}
