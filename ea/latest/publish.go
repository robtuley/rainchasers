package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"time"
)

func publish(
	projectID string,
	topicName string,
	readings map[string]gauge.Reading,
	stations map[string]gauge.Station,
) error {
	throttle := time.NewTicker(time.Second / maxPublishPerSecond)
	topic, err := queue.New(projectID, topicName)
	if err != nil {
		return err
	}

	for id, r := range readings {
		s, ok := stations[id]
		if !ok {
			continue
		}

		<-throttle.C
		err := topic.Publish(&gauge.Snapshot{
			Station:  s,
			Readings: []gauge.Reading{r},
		})
		if err != nil {
			return err
		}
	}

	throttle.Stop()
	topic.Stop()

	return nil
}
