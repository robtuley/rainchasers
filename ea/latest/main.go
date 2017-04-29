package main

import (
	"errors"
	"github.com/rainchasers/com.rainchasers.gauge/ea/discover"
	"github.com/rainchasers/report"
	"os"
	"strconv"
	"time"
)

// Responds to environment variables:
//   UPDATE_EVERY_X_SECONDS (default 15*60)
//   SHUTDOWN_AFTER_X_SECONDS (default 24*60*60)
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
//
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	// parse env vars
	updatePeriodSeconds, err := strconv.Atoi(os.Getenv("UPDATE_EVERY_X_SECONDS"))
	if err != nil {
		updatePeriodSeconds = 15 * 60
	}
	shutdownDeadline, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		shutdownDeadline = 7 * 24 * 60 * 60
	}
	shutdownC := time.NewTimer(time.Second * time.Duration(shutdownDeadline)).C
	projectId := os.Getenv("PROJECT_ID")
	isValidating := projectId == ""
	topicName := os.Getenv("PUBSUB_TOPIC")

	// setup telemetry
	log := report.New(report.Data{"service": "ea.latest", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	log.RuntimeStatEvery("runtime", 5*time.Minute)
	defer log.Stop()
	log.Info("daemon.start", report.Data{
		"update_period":     updatePeriodSeconds,
		"shutdown_deadline": shutdownDeadline,
		"project_id":        projectId,
		"pubsub_topic":      topicName,
	})

	// discover EA gauging stations
	stations, err := discover.Stations()
	if err != nil {
		<-log.Action("discovered.fail", report.Data{"error": err.Error()})
		return err
	}
	log.Info("discovered.ok", report.Data{"count": len(stations)})

	// periodically get latest updates
	ticker := time.NewTicker(time.Second * time.Duration(updatePeriodSeconds))

updateLoop:
	for {
		tick := log.Tick()
		updates, err := update()
		if err != nil {
			<-log.Action("updated.fail", report.Data{"error": err.Error()})
			return err
		}
		log.Tock(tick, "updated.ok", report.Data{"count": len(updates)})

		if !isValidating {
			tick = log.Tick()
			err = publish(projectId, topicName, updates, stations)
			if err != nil {
				<-log.Action("published.fail", report.Data{"error": err.Error()})
				return err
			}
			log.Tock(tick, "published.ok", report.Data{"count": len(updates)})
		}

		select {
		case <-ticker.C:
		case <-shutdownC:
			break updateLoop
		}
	}
	ticker.Stop()

	// validate log stream on shutdown if required
	if isValidating {
		if log.Count("discovered.ok") != 1 {
			return errors.New("discovered.ok event expected but not present")
		}
		if log.Count("updated.ok") < 1 {
			return errors.New("updated.ok event expected but not present")
		}
	}

	return nil
}
