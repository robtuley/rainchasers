package main

import (
	"encoding/json"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/request"
)

type readingJSON struct {
	Items []struct {
		Measure      string          `json:"measure"`
		DateTime     time.Time       `json:"dateTime"`
		ValueRawJSON json.RawMessage `json:"value"`
	} `json:"items"`
}

func update() (map[string]gauge.Reading, error) {
	url := "http://environment.data.gov.uk/flood-monitoring/data/readings?latest"
	readings := make(map[string]gauge.Reading)

	resp, err := request.JSON(url)
	if err != nil {
		return readings, err
	}
	defer resp.Body.Close()

	r := readingJSON{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return readings, err
	}

	for _, item := range r.Items {
		// the 'value' keys should be a float, but instances exist of arrays
		// so we do a conditional parse and simply dump those that can't match.
		value, err := request.ParseFloat(item.ValueRawJSON)
		if err != nil {
			continue
		}

		readings[item.Measure] = gauge.Reading{
			EventTime: item.DateTime,
			Value:     value,
		}
	}

	return readings, nil
}
