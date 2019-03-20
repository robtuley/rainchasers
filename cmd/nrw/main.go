package main

import (
	"bytes"
	"context"
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
//   NRW_API_KEY (no default)
func main() {
	cfg := config{
		ProjectID:                os.Getenv("PROJECT_ID"),
		TopicName:                os.Getenv("PUBSUB_TOPIC"),
		APIKey:                   os.Getenv("NRW_API_KEY"),
		RefreshPeriodInSeconds:   15 * 60,
		MaxPublishPerSecond:      30,
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
	APIKey                   string
	RefreshPeriodInSeconds   int
	MaxPublishPerSecond      int
	ExitAfterXConsecutiveErr int
}

func (cfg config) run(ctx context.Context, d *daemon.Supervisor) error {
	// open connection to pubsub
	topic, span := queue.New(ctx, cfg.ProjectID, cfg.TopicName)
	d.Trace(span)
	if err := span.Err(); err != nil {
		return err
	}
	defer topic.Stop()

	nConsecutiveErr := 0
pollLoop:
	for {
		err := func() error {
			// get recent stations & readings
			var snapshots []gauge.Snapshot
			var err error
			var span report.Span
			if cfg.APIKey == "" {
				snapshots, err = parseRecent(bytes.NewBufferString(jsonResponseFromAPI))
			} else {
				snapshots, span = recent(ctx, cfg.APIKey)
				d.Trace(span)
				err = span.Err()
			}
			if err != nil {
				return err
			}

			// calculate update rate to refresh on schedule
			every := cfg.durationBetweenPublish(len(snapshots))
			ticker := time.NewTicker(every)
			defer ticker.Stop()

			// publish snapshots
			for _, s := range snapshots {
				span := topic.Publish(ctx, &s)
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

			return nil
		}()

		// exit loop if shutdown
		select {
		case <-ctx.Done():
			break pollLoop
		default:
		}

		// track consecutive errors
		if err != nil {
			nConsecutiveErr++
			if nConsecutiveErr >= cfg.ExitAfterXConsecutiveErr {
				// ignore a few isolated errors, but if
				// many consecutive bubble up to restart
				return err
			}
		} else {
			nConsecutiveErr = 0
		}
	}

	return nil
}

func (cfg config) durationBetweenPublish(total int) time.Duration {
	ms := cfg.RefreshPeriodInSeconds * 1000 / total
	min := 1
	if cfg.MaxPublishPerSecond > 0 {
		min = 1000 / cfg.MaxPublishPerSecond
	}
	if ms < min {
		ms = min
	}
	return time.Millisecond * time.Duration(ms)
}
