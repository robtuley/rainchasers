package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"time"
)

func publish(
	projectID string,
	topicName string,
	updates []gauge.SnapshotUpdate,
) error {
	// groups updates per metric ID
	var updateMap map[string][]gauge.SnapshotUpdate
	for _, u := range updates {
		updateMap[u.MetricID] = append(updateMap[u.MetricID], u)
	}

	throttle := time.NewTicker(time.Second / maxPublishPerSecond)
	ctx, err := gauge.NewPubSubContext(projectID, topicName)
	_ = ctx
	if err != nil {
		return err
	}

	for id, snaps := range updateMap {
		_ = id
		_ = snaps
		<-throttle.C
	}
	throttle.Stop()

	return nil
}
