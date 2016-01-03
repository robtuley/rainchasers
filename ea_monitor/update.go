package main

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

type readingJson struct {
	Items []struct {
		Measure      string          `json:"measure"`
		DateTime     time.Time       `json:"dateTime"`
		ValueRawJson json.RawMessage `json:"value"`
		ValueParsed  float32
	} `json:"items"`
}

func requestLatestReadings(updateC chan gauge.SnapshotUpdate) {
	const BATCHSIZE = 100

	lastBatchSize := BATCHSIZE
	currentOffset := 0
	baseUrl := "http://environment.data.gov.uk/flood-monitoring/data/readings" +
		"?latest&_limit=" + strconv.Itoa(BATCHSIZE)

	for lastBatchSize > 0 {
		waitOnApiRequestSchedule()

		url := baseUrl + "&_offset=" + strconv.Itoa(currentOffset)
		currentOffset = currentOffset + BATCHSIZE
		tick := report.Tick()

		resp, err := doJsonRequest(url)
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
			// the 'value' keys should be a float, but instances exist of arrays
			// so we do a conditional parse and simply dump those that can't match.
			err := json.Unmarshal(item.ValueRawJson, &item.ValueParsed)
			if err != nil {
				report.Info("update.corrupt.value",
					report.Data{"url": item.Measure, "json": item.ValueRawJson})
				continue
			}

			updateC <- gauge.SnapshotUpdate{
				item.Measure,
				item.DateTime,
				item.ValueParsed,
			}
		}
	}

}
