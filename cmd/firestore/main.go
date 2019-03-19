package main

import (
	"context"
	"os"
	"sync/atomic"
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
	d.Run(context.Background(), app.LogAndResetCountStats)
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
	// CountReceived snapshots & readings, access using sync.atomic
	CountSnapshotsReceived uint64
	CountReadingsReceived  uint64
}

func (c *cache) PollForRiverCatalogueChanges(ctx context.Context, d *daemon.Supervisor) error {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return nil
		}

		isChanged, err := c.Rivers.Update(ctx, d)
		if err != nil {
			return err
		}
		if isChanged {
			err := c.OnRiverCatalogueChange(ctx, d)
			if err != nil {
				return err
			}
		}
	}
}

func (c *cache) OnRiverCatalogueChange(ctx context.Context, d *daemon.Supervisor) error {
	// TODO: update firebase rivers
	return nil
}

func (c *cache) LogAndResetCountStats(ctx context.Context, d *daemon.Supervisor) error {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return nil
		}

		nSnaps := atomic.SwapUint64(&c.CountSnapshotsReceived, 0)
		nReadings := atomic.SwapUint64(&c.CountReadingsReceived, 0)
		d.Info("snapshot.received", report.Data{
			"count":         nSnaps,
			"reading_count": nReadings,
		})
	}
}

func (c *cache) SubscribeToSnapshotUpdates(ctx context.Context, d *daemon.Supervisor) error {
	topic, err := queue.New(ctx, d, c.ProjectID, c.TopicName)
	if err != nil {
		return err
	}
	defer topic.Stop()

	return topic.Subscribe(ctx, d, "", c.OnReceivedSnapshot)
}

func (c *cache) OnReceivedSnapshot(ctx context.Context, d *daemon.Supervisor, s *gauge.Snapshot) error {
	atomic.AddUint64(&c.CountSnapshotsReceived, 1)
	atomic.AddUint64(&c.CountReadingsReceived, uint64(len(s.Readings)))

	// only return error if want message redelivered, otherwise deal with it locally
	return nil
}
