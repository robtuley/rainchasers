package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"time"
)

func publish(
	projectID string,
	topicName string,
	updates map[string][]gauge.SnapshotUpdate,
) error {
	throttle := time.NewTicker(time.Second / maxPublishPerSecond)
	defer throttle.Stop()

	_, err := gauge.NewPubSubContext(projectID, topicName)
	if err != nil {
		return err
	}

	for url, snaps := range updates {
		_, err := gauge.EncodeSnapshotUpdates(url, snaps)
		if err != nil {
			return err
		}
		<-throttle.C
	}

	return nil
}
