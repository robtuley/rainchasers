package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/ea"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/internal/queue"
)

// Responds to environment variables:
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
func main() {
	d := daemon.New("ea")
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
	isDryRun := projectID == ""
	refreshPeriodInSeconds := 5 * 60

	// discover EA gauging stations
	stations, err := ea.Discover(ctx, d)
	if err != nil {
		return err
	}

	// if dry run, shorten running model
	if isDryRun {
		n := 0
		for k := range stations {
			if n >= 3 {
				delete(stations, k)
			}
			n++
		}
		go func() {
			<-time.After(10 * time.Second)
			d.Close()
		}()
		refreshPeriodInSeconds = 3
	}

	nConsecutiveErr := 0
updateLoop:
	for {
		err := func(ctx context.Context) error {
			// get all recent readings
			readings, err := ea.Recent(ctx, d)
			if err != nil {
				return err
			}

			// open connection to pubsub
			topic, err := queue.New(ctx, d, projectID, topicName)
			if err != nil {
				return err
			}
			defer topic.Stop()

			// calculate tick rate to spread readings publish over
			// the refresh period
			every := calculateRate(len(readings), refreshPeriodInSeconds)
			ticker := time.NewTicker(every)
			defer ticker.Stop()

			// publish readings
			for id, r := range readings {
				s, ok := stations[id]
				if !ok {
					continue
				}

				err := topic.Publish(ctx, d, &gauge.Snapshot{
					Station:  s,
					Readings: []gauge.Reading{r},
				})
				if err != nil {
					return err
				}

				select {
				case <-ticker.C:
				case <-ctx.Done():
					// exit early on shutdown
					return nil
				}
			}

			return nil
		}(ctx)

		if err != nil {
			nConsecutiveErr++
			if nConsecutiveErr > 3 {
				// ignore a few isolated errors, but if
				// many consecutive bubble up to restart
				return err
			}
		} else {
			nConsecutiveErr = 0
		}

		// break loop on shutdown signal
		select {
		case <-d.Done():
			break updateLoop
		default:
		}
	}

	// validate log stream on shutdown
	if d.Count("snapshot.published") < 1 {
		return errors.New("No snapshot.published events")
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
