package main

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/robtuley/report"
)

type readingJson struct {
	Items []struct {
		Measure  string    `json:"measure"`
		DateTime time.Time `json:"dateTime"`
		Value    float32   `json:"value"`
	} `json:"items"`
}

func requestLatestReadings(updateC chan *SnapshotUpdate) {
	batchSize := 100
	lastBatchSize := batchSize
	currentOffset := 0
	baseUrl := "http://environment.data.gov.uk/flood-monitoring/data/readings" +
		"?latest&_limit=" + strconv.Itoa(batchSize)

	for lastBatchSize == batchSize {
		waitOnApiRequestSchedule()

		url := baseUrl + "&_offset=" + strconv.Itoa(currentOffset)
		currentOffset = currentOffset + batchSize
		tick := report.Tick()

		err, resp := doJsonRequest(url)
		if err != nil {
			report.Action("update.request.error", report.Data{"url": url, "error": err.Error()})
			continue
		}

		r := readingJson{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&r)
		resp.Body.Close()
		if err != nil {
			report.Action("update.decode.error", report.Data{"url": url, "error": err.Error()})
			continue
		}

		lastBatchSize = len(r.Items)
		report.Tock(tick, "update.response", report.Data{
			"count": lastBatchSize,
			"url":   url,
		})

		for _, item := range r.Items {
			updateC <- &SnapshotUpdate{
				item.Measure,
				item.DateTime,
				item.Value,
			}
		}
	}

}
