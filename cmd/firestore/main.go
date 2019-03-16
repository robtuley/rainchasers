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

// Env:
//   PROJECT_ID (no default, blank indicates self-test mode)
//   PUBSUB_TOPIC (no default)
func main() {
	cfg := config{
		ProjectID: os.Getenv("PROJECT_ID"),
		TopicName: os.Getenv("PUBSUB_TOPIC"),
	}

	d := daemon.New("firestore")
	go d.Run(context.Background(), cfg.run)
	select {
	case <-time.After(24 * time.Hour):
	case <-d.Done():
	}
	d.Close()

	if err := d.Err(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

type config struct {
	ProjectID string
	TopicName string
}

func (cfg config) run(ctx context.Context, d *daemon.Supervisor) error {
	// load river catalogue
	url := "https://app.rainchasers.com/catalogue.json"
	rivers, err := NewRiverCache(ctx, d, url)
	if err != nil {
		return err
	}
	// TODO: update firebase rivers here

	// poll for river content updates
	go d.Run(ctx, func(ctx context.Context, d *daemon.Supervisor) error {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return nil
			}

			isChanged, err := rivers.Update(ctx, d)
			if err != nil {
				return err
			}
			if isChanged {
				// TODO: update firebase rivers here
			}
		}
	})

	// subscribe to gauge snapshot topic to populate firebase
	var counter uint64
	go d.Run(ctx, func(ctx context.Context, d *daemon.Supervisor) error {
		if cfg.ProjectID == "" {
			return nil
		}

		topic, err := queue.New(ctx, d, cfg.ProjectID, cfg.TopicName)
		if err != nil {
			return err
		}
		defer topic.Stop()

		return topic.Subscribe(ctx, d, "", func(d *daemon.Supervisor, s *gauge.Snapshot, err error) {
			if err != nil {
				d.Action("msg.failed", report.Data{"error": err.Error()})
			} else {
				atomic.AddUint64(&counter, 1)
				// do something with firebase here
			}
		})
	})

	// log updates received ok every 30s
	go d.Run(ctx, func(ctx context.Context, d *daemon.Supervisor) error {
		ticker := time.NewTicker(time.Second * 30)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return nil
			}

			d.Info("received.ok", report.Data{
				"count": atomic.LoadUint64(&counter),
			})
			atomic.StoreUint64(&counter, 0)
		}
	})

	return nil
}
