package main

import (
	"errors"
	"os"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/ea/discover"
	"github.com/rainchasers/report"
)

// Responds to environment variables:
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	// parse env vars
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	// setup config, blank project ID switches to validation run
	updatePeriod := 15 * 60 * time.Second
	shutdownDeadline := 7 * 24 * time.Hour
	if projectID == "" {
		updatePeriod = 5 * time.Second
		shutdownDeadline = 15 * time.Second
	}
	shutdownC := time.NewTimer(shutdownDeadline).C

	// setup telemetry
	log := report.New(report.StdOutJSON(), report.Data{"service": "ea.latest", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	log.RuntimeStatEvery("runtime", 5*time.Minute)
	defer log.Stop()
	log.Info("daemon.start", report.Data{
		"project_id":   projectID,
		"pubsub_topic": topicName,
	})

	// discover EA gauging stations
	stations, err := discover.Stations()
	if err != nil {
		<-log.Action("discovered.fail", report.Data{"error": err.Error()})
		return err
	}
	if projectID == "" { // limit map size to 5 stations for validation
		n := 0
		for k := range stations {
			if n >= 5 {
				delete(stations, k)
			}
			n++
		}
	}
	log.Info("discovered.ok", report.Data{"count": len(stations)})

	// periodically get latest updates
	ticker := time.NewTicker(updatePeriod)

updateLoop:
	for {
		tick := log.Tick()
		updates, err := update()
		if err != nil {
			<-log.Action("updated.fail", report.Data{"error": err.Error()})
			return err
		}
		log.Tock(tick, "updated.ok", report.Data{"count": len(updates)})

		tick = log.Tick()
		if err = publish(projectID, topicName, updates, stations); err != nil {
			<-log.Action("published.fail", report.Data{"error": err.Error()})
			return err
		}
		log.Tock(tick, "published.ok", report.Data{"count": len(updates)})

		select {
		case <-ticker.C:
		case <-shutdownC:
			break updateLoop
		}
	}
	ticker.Stop()

	// validate log stream on shutdown if required
	if log.Count("discovered.ok") != 1 {
		return errors.New("discovered.ok event expected but not present")
	}
	if log.Count("updated.ok") < 1 {
		return errors.New("updated.ok event expected but not present")
	}
	if err := log.LastError(); err != nil {
		return err
	}

	return nil
}
