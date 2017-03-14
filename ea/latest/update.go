package main

import (
	"encoding/json"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"time"
)

type readingJson struct {
	Items []struct {
		Measure      string          `json:"measure"`
		DateTime     time.Time       `json:"dateTime"`
		ValueRawJson json.RawMessage `json:"value"`
	} `json:"items"`
}

func update() ([]gauge.SnapshotUpdate, error) {
	url := "http://environment.data.gov.uk/flood-monitoring/data/readings?latest"
	var updates []gauge.SnapshotUpdate

	resp, err := doJsonRequest(url)
	if err != nil {
		return updates, err
	}
	defer resp.Body.Close()

	r := readingJson{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return updates, err
	}

	for _, item := range r.Items {
		// the 'value' keys should be a float, but instances exist of arrays
		// so we do a conditional parse and simply dump those that can't match.
		value, err := parseFloatFromScalarOrArray(item.ValueRawJson)
		if err != nil {
			continue
		}

		updates = append(updates, gauge.SnapshotUpdate{
			MetricID: gauge.CalculateMetricID(item.Measure),
			DateTime: item.DateTime,
			Value:    value,
		})
	}

	return updates, nil
}
