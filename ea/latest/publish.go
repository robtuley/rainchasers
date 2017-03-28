package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"time"
)

func publish(
	projectID string,
	topicName string,
	updates map[string]gauge.SnapshotUpdate,
	refSnapshots map[string]gauge.Snapshot,
) error {
	throttle := time.NewTicker(time.Second / maxPublishPerSecond)
	topic, err := gauge.NewTopic(projectID, topicName)
	if err != nil {
		return err
	}

	for id, u := range updates {
		s, ok := refSnapshots[id]
		if !ok {
			continue
		}

		<-throttle.C
		err := gauge.Publish(topic, s.Apply(u))
		if err != nil {
			return err
		}
	}
	throttle.Stop()
	topic.Stop()

	return nil
}
