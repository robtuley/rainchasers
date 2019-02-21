package main

import (
	"errors"
	"os"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/ea/discover"
	"github.com/rainchasers/report"
)

const maxPublishPerSecond = 20

// Responds to environment variables:
//   DATE (defaults to yesterday)
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	var err error

	// parse env vars
	requestedDay := os.Getenv("DATE")
	var day time.Time
	if len(requestedDay) == 0 {
		day = time.Now().AddDate(0, 0, -2)
	} else {
		day, err = time.Parse("2006-01-02", requestedDay)
		if err != nil {
			return err
		}
	}
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	// setup telemetry
	log := report.New(report.StdOutJSON(), report.Data{"service": "ea.latest", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	log.RuntimeStatEvery("runtime", 10*time.Second)
	defer log.Stop()
	log.Info("daemon.start", report.Data{
		"day":          day.Format("2006-01-02"),
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

	// download & parse CSV data
	tick := log.Tick()
	readings, err := download(day)
	if err != nil {
		<-log.Action("download.fail", report.Data{"error": err.Error()})
		return err
	}
	log.Tock(tick, "download.ok", report.Data{"count": len(readings)})

	// publish historical data
	tick = log.Tick()
	if err := publish(projectID, topicName, stations, readings); err != nil {
		<-log.Action("publish.fail", report.Data{"error": err.Error()})
		return err
	}
	log.Tock(tick, "publish.ok", report.Data{})

	// validate log stream on shutdown if required
	if log.Count("discovered.ok") != 1 {
		return errors.New("discovered.ok event expected but not present")
	}
	if log.Count("download.ok") != 1 {
		return errors.New("download.ok event expected but not present")
	}
	if err := log.LastError(); err != nil {
		return err
	}

	return nil
}
