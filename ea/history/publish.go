package main

import (
	"context"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"time"
)

func publish(
	projectID string,
	topicName string,
	stations map[string]gauge.Station,
	readings map[string][]gauge.Reading,
) error {
	throttle := time.NewTicker(time.Second / maxPublishPerSecond)
	defer throttle.Stop()

	topic, err := queue.New(context.Background(), projectID, topicName)
	if err != nil {
		return err
	}
	defer topic.Stop()

	for id, r := range readings {
		s, ok := stations[id]
		if !ok {
			continue
		}

		<-throttle.C
		err := topic.Publish(context.Background(), &gauge.Snapshot{
			Station:  s,
			Readings: r,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
