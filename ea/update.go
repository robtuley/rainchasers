package main

import (
	"encoding/json"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

type readingJson struct {
	Items []struct {
		Measure      string          `json:"measure"`
		DateTime     time.Time       `json:"dateTime"`
		ValueRawJson json.RawMessage `json:"value"`
	} `json:"items"`
}

func requestLatestReadings(updateC chan gauge.SnapshotUpdate) {
	url := "http://environment.data.gov.uk/flood-monitoring/data/readings?latest"
	waitOnApiRequestSchedule()
	tick := report.Tick()

	resp, err := doJsonRequest(url)
	if err != nil {
		report.Action("update.request.error", report.Data{"url": url, "error": err.Error()})
		return
	}

	r := readingJson{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&r)
	resp.Body.Close()
	if err != nil {
		report.Action("update.decode.error", report.Data{"url": url, "error": err.Error()})
		return
	}

	report.Tock(tick, "update.response", report.Data{
		"count": len(r.Items),
		"url":   url,
	})

	for _, item := range r.Items {
		// the 'value' keys should be a float, but instances exist of arrays
		// so we do a conditional parse and simply dump those that can't match.
		value, err := parseFloatFromScalarOrArray(item.ValueRawJson)
		if err != nil {
			report.Info("update.value.missing",
				report.Data{"url": item.Measure, "json": item.ValueRawJson})
			continue
		}

		u := gauge.SnapshotUpdate{
			MetricID: gauge.CalculateMetricID(item.Measure),
			DateTime: item.DateTime,
			Value:    value,
		}

		updateC <- u
	}
}
