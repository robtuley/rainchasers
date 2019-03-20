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

const maxPublishPerSecond = 20

// Responds to environment variables:
//   DATE (defaults to yesterday)
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
func main() {
	d := daemon.New("gauge")
	d.Run(context.Background(), run)
	d.CloseAfter(4 * time.Hour)
	d.Wait()
	if err := d.Err(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run(ctx context.Context, d *daemon.Supervisor) error {
	// parse env vars
	requestedDay := os.Getenv("DATE")
	var day time.Time
	if len(requestedDay) == 0 {
		day = time.Now().AddDate(0, 0, -2)
	} else {
		var err error
		day, err = time.Parse("2006-01-02", requestedDay)
		if err != nil {
			return err
		}
	}
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	isDryRun := projectID == ""

	// discover EA gauging stations
	stations, dSpan := ea.Discover(ctx)
	if err := dSpan.Err(); err != nil {
		d.Trace(dSpan)
		return err
	}

	// if dry run shorten the run
	if isDryRun {
		n := 0
		for k := range stations {
			if n >= 3 {
				delete(stations, k)
			}
			n++
		}
	}

	// get all readings on requested day
	readings, rSpan := ea.Day(ctx, day)
	if err := rSpan.Err(); err != nil {
		d.Trace(dSpan.FollowedBy(rSpan))
		return err
	}

	// open connection to pubsub
	topic, cSpan := queue.New(ctx, projectID, topicName)
	d.Trace(dSpan.FollowedBy(rSpan).FollowedBy(cSpan))
	if err := cSpan.Err(); err != nil {
		return err
	}
	defer topic.Stop()

	// publish readings
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for id, r := range readings {
		s, ok := stations[id]
		if !ok {
			continue
		}

		span := topic.Publish(ctx, &gauge.Snapshot{
			Station:  s,
			Readings: r,
		})
		d.Trace(span)
		if err := span.Err(); err != nil {
			return err
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			// exit early on shutdown
			return nil
		}
	}

	// validate log stream on shutdown
	if d.Count("snapshot.published") < 1 {
		return errors.New("No snapshot.published events")
	}

	return nil
}
