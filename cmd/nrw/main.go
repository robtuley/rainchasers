package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/internal/queue"
)

const refreshPeriodInSeconds = 15 * 60

// Responds to environment variables:
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
//   API_KEY (no default)
func main() {
	d := daemon.New("nrw")
	go d.Run(context.Background(), run)

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

func run(ctx context.Context, d *daemon.Supervisor) error {
	// parse env vars
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	apiKey := os.Getenv("API_KEY")
	isDryRun := projectID == ""
	if !isDryRun && apiKey == "" {
		return errors.New("No API_KEY env var provided")
	}

	// open connection to pubsub
	topic, err := queue.New(ctx, d, projectID, topicName)
	if err != nil {
		return err
	}
	defer topic.Stop()

	nConsecutiveErr := 0
	for {
		err := func(ctx context.Context) error {
			// get recent stations & readings
			snapshots := make([]gauge.Snapshot, 0)

			// calculate update rate to refresh on schedule
			every := calculateRate(len(snapshots), refreshPeriodInSeconds)
			ticker := time.NewTicker(every)
			defer ticker.Stop()

		pubLoop:
			for _, s := range snapshots {
				err := topic.Publish(ctx, d, &s)
				if err != nil {
					return err
				}

				select {
				case <-ticker.C:
				case <-ctx.Done():
					break pubLoop
				}
			}

			return nil
		}(ctx)

		if err != nil {
			nConsecutiveErr++
			if nConsecutiveErr > 2 {
				// ignore a few isolated errors, but if
				// many consecutive bubble up to restart
				return err
			}
		} else {
			nConsecutiveErr = 0
		}
	}

	// validate log stream
	if d.Count("sepa.discover") != 1 {
		return errors.New("sepa.discover event expected but not present")
	}
	if d.Count("sepa.updated") < 1 {
		return errors.New("sepa.updated event expected but not present")
	}

	return nil
}

func calculateRate(n int, seconds int) time.Duration {
	maxPerSecond := 50
	ms := seconds * 1000 / n
	min := 1000 / maxPerSecond
	if ms < min {
		ms = min
	}
	return time.Millisecond * time.Duration(ms)
}
