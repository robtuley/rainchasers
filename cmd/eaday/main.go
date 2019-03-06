package main

import (
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
	d := daemon.New("ea", 24*time.Hour)
	d.Run(run)
	d.Close()

	if err := d.Err(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run(d *daemon.Supervisor) error {
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
	stations, err := ea.Discover(d)
	if err != nil {
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
	readings, err := ea.Day(d, day)
	if err != nil {
		return err
	}

	// open connection to pubsub
	topic, err := queue.New(d, projectID, topicName)
	if err != nil {
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

		err := topic.Publish(d, &gauge.Snapshot{
			Station:  s,
			Readings: r,
		})
		if err != nil {
			return err
		}

		select {
		case <-ticker.C:
		case <-d.Context().Done():
			// exit early on shutdown
			return nil
		}
	}

	// validate log stream on shutdown
	if d.Count("snapshot.published") < 1 {
		return errors.New("No snapshot.published events")
	}
	if err := d.Err(); err != nil {
		return err
	}

	return nil
}
