package ea

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

// Recent fetches recent EA readings
func Recent(ctx context.Context) (map[string]gauge.Reading, report.Span) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	span := report.StartSpan("ea.recent").Field("url", recentReadingURL)
	readings := make(map[string]gauge.Reading)

	resp, err := daemon.JSON(ctx, recentReadingURL)
	if err != nil {
		return readings, span.End(err)
	}
	defer resp.Body.Close()

	r := readingJSON{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return readings, span.End(err)
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

	span = span.Field("readings_count", len(readings))
	return readings, span.End()
}
