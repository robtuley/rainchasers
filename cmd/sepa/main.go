package main

import (
	"context"
	"os"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/internal/queue"
)

// Responds to environment variables:
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
func main() {
	cfg := config{
		ProjectID:                os.Getenv("PROJECT_ID"),
		TopicName:                os.Getenv("PUBSUB_TOPIC"),
		RefreshPeriodInSeconds:   15 * 60,
		ExitAfterXConsecutiveErr: 3,
	}

	d := daemon.New("gauge")
	d.Run(context.Background(), cfg.run)
	d.CloseAfter(24 * time.Hour)

	d.Wait()
	if err := d.Err(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

type config struct {
	ProjectID                string
	TopicName                string
	RefreshPeriodInSeconds   int
	ExitAfterXConsecutiveErr int
}

func (cfg config) run(ctx context.Context, d *daemon.Supervisor) error {
	// discover SEPA gauging stations
	stations, dSpan := discover(ctx)
	if err := dSpan.Err(); err != nil {
		d.Trace(dSpan)
		return err
	}

	// calculate update rate to refresh on schedule
	every := cfg.durationBetweenPublish(len(stations))
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	// open connection to pubsub
	topic, qSpan := queue.New(ctx, cfg.ProjectID, cfg.TopicName)
	d.Trace(dSpan.FollowedBy(qSpan))
	if err := qSpan.Err(); err != nil {
		return err
	}
	defer topic.Stop()

	// get readings & publish them to pubsub
	nConsecutiveErr := 0
	n := 0
updateLoop:
	for {
		i := n % len(stations)

		readings, span := getReadings(ctx, stations[i].DataURL)
		if err := span.Err(); err == nil {
			span = span.FollowedBy(topic.Publish(ctx, &gauge.Snapshot{
				Station:       stations[i],
				Readings:      readings,
				CorrelationID: span.TraceID(),
				CausationID:   span.SpanID(),
			}))
		}

		d.Trace(span)

		if err := span.Err(); err != nil {
			nConsecutiveErr++
			if nConsecutiveErr >= cfg.ExitAfterXConsecutiveErr {
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
		case <-ctx.Done():
			break updateLoop
		}
	}

	return nil
}

func (cfg config) durationBetweenPublish(total int) time.Duration {
	ms := cfg.RefreshPeriodInSeconds * 1000 / total
	min := 1500 // SEPA rate limiter ~ 1 req/per second
	if ms < min {
		ms = min
	}
	return time.Millisecond * time.Duration(ms)
}
