package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"time"
)

func publish(
	projectID string,
	topicName string,
	readings map[string][]gauge.Reading,
) error {
	throttle := time.NewTicker(time.Second / maxPublishPerSecond)
	defer throttle.Stop()

	_ = projectID
	_ = topicName

	for url, data := range readings {
		_ = url
		_ = data
		<-throttle.C
	}

	return nil
}
