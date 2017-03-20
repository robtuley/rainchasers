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
	ctx, err := gauge.NewPubSubContext(projectID, topicName)
	if err != nil {
		return err
	}

	for id, u := range updates {
		s, ok := refSnapshots[id]
		if !ok {
			continue
		}

		<-throttle.C
		err := gauge.Publish(ctx, s.Apply(u))
		if err != nil {
			return err
		}
	}
	throttle.Stop()

	return nil
}
