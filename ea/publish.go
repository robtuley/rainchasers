package main

import (
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

func logSnapshotsFromChannel(label string, snapC <-chan gauge.Snapshot) {
	for s := range snapC {
		encoded, err := gauge.EncodeSnapshot(s)
		if err != nil {
			report.Action("snapshot.encode.error", report.Data{"error": err.Error(), "snapshot": s})
			continue
		}
		_, err = gauge.DecodeSnapshot(encoded)
		if err != nil {
			report.Action("snapshot.decode.error", report.Data{"error": err.Error(), "snapshot": s})
			continue
		}
		report.Info(label, report.Data{"snapshot": s})
	}
}

func publishSnapshotsFromChannels(
	projectId string, latestTopicName string, historyTopicName string,
	latestSnapC <-chan gauge.Snapshot, historySnapC <-chan gauge.Snapshot,
) error {

	latestCtx, err := gauge.NewPubSubContext(projectId, latestTopicName)
	if err != nil {
		return err
	}
	historyCtx, err := gauge.NewPubSubContext(projectId, historyTopicName)
	if err != nil {
		return err
	}

	go func() {
		tickC := time.Tick(time.Second * 10)
		nLatest := 0
		nHistory := 0
		for {
			select {
			case s, is_ok := <-latestSnapC:
				if !is_ok {
					break
				}
				err := gauge.Publish(latestCtx, s)
				nLatest = nLatest + 1
				if err != nil {
					report.Action("pubsub.publish.latest", report.Data{
						"error": err.Error(),
					})
				}
			case s, is_ok := <-historySnapC:
				if !is_ok {
					break
				}
				err := gauge.Publish(historyCtx, s)
				nHistory = nHistory + 1
				if err != nil {
					report.Action("pubsub.publish.history", report.Data{
						"error": err.Error(),
					})
				}
			case <-tickC:
				report.Info("pubsub.publish.ok", report.Data{"latest": nLatest, "history": nHistory})
				nLatest = 0
				nHistory = 0
			}
		}
	}()

	return nil
}
