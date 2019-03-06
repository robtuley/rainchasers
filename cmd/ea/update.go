package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/report"
)

type readingJSON struct {
	Items []struct {
		Measure      string          `json:"measure"`
		DateTime     time.Time       `json:"dateTime"`
		ValueRawJSON json.RawMessage `json:"value"`
	} `json:"items"`
}

const recentReadingURL = "http://environment.data.gov.uk/flood-monitoring/data/readings?latest"

func update(d *daemon.Supervisor) (rd map[string]gauge.Reading, err error) {
	ctx, cancel := context.WithTimeout(d.Context(), 60*time.Second)
	ctx = d.StartSpan(ctx, "recent.readings")
	defer func() {
		d.EndSpan(ctx, err, report.Data{
			"count": len(rd),
			"url":   recentReadingURL,
		})
		cancel()
	}()

	readings := make(map[string]gauge.Reading)

	resp, err := daemon.JSON(ctx, recentReadingURL)
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
		value, err := daemon.ParseFloat(item.ValueRawJSON)
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
