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

// Responds to environment variables:
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
func main() {
	d := daemon.New("sepa")
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

	// discover SEPA gauging stations
	stations, err := discover(ctx, d)
	if err != nil {
		return err
	}

	// calculate update rate to refresh on schedule
	refreshPeriodInSeconds := 15 * 60
	every := calculateRate(len(stations), refreshPeriodInSeconds)
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	// if dry run, shorten running model
	if isDryRun {
		stations = stations[0:3]
		go func() {
			<-time.After(10 * time.Second)
			d.Close()
		}()
		ticker = time.NewTicker(2 * time.Second)
		defer ticker.Stop()
	}

	// open connection to pubsub
	topic, err := queue.New(ctx, d, projectID, topicName)
	if err != nil {
		return err
	}
	defer topic.Stop()

	// get readings & publish them to pubsub
	nConsecutiveErr := 0
	n := 0
updateLoop:
	for {
		i := n % len(stations)

		err := func(ctx context.Context) (err error) {
			ctx = d.Trace(ctx)
			ctx = d.StartSpan(ctx, "sepa.updated")
			defer func() {
				d.EndSpan(ctx, err, report.Data{
					"station": stations[i].UUID(),
				})
			}()

			readings, err := getReadings(ctx, d, stations[i].DataURL)
			if err != nil {
				return err
			}

			return topic.Publish(ctx, d, &gauge.Snapshot{
				Station:  stations[i],
				Readings: readings,
			})
		}(ctx)

		if err != nil {
			nConsecutiveErr++
			if nConsecutiveErr > 5 {
				// ignore a few isolated errors, but if
				// many consecutive bubble up to restart
				return err
			}
		} else {
			nConsecutiveErr = 0
		}

		n++
		select {
		case <-ticker.C:
		case <-d.Done():
			break updateLoop
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
	ms := seconds * 1000 / n
	min := 1500 // SEPA rate limiter ~ 1 req/per second
	if ms < min {
		ms = min
	}
	return time.Millisecond * time.Duration(ms)
}
