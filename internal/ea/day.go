package ea

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/robtuley/rainchasers/internal/daemon"
	"github.com/robtuley/rainchasers/internal/gauge"
	"github.com/robtuley/report"
)

// Day downloads all the measurements on specified day
func Day(ctx context.Context, day time.Time) (map[string][]gauge.Reading, report.Span) {
	url := "http://environment.data.gov.uk/flood-monitoring/archive/readings-" + day.Format("2006-01-02") + ".csv"
	readings := make(map[string][]gauge.Reading)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	span := report.StartSpan("ea.day").Field("url", url)

	resp, err := daemon.CSV(ctx, url)
	if err != nil {
		return readings, span.End(err)
	}
	defer resp.Body.Close()

	csv := csv.NewReader(resp.Body)
	isFirst := true

ReadCSV:
	for {
		r, err := csv.Read()

		if err == io.EOF || err == io.ErrUnexpectedEOF || err == io.ErrClosedPipe {
			break ReadCSV
		}
		// some corrupt reading values appear as 1.23|4.56 so
		// we simply skip these as known errors.
		if len(r) == 3 {
			if strings.Contains(r[2], "|") {
				continue
			}
		}
		if err != nil {
			return readings, span.End(err)
		}
		if isFirst {
			isFirst = false
			continue
		}

		url, s, err := csvRecordToReading(r)
		if err != nil {
			return readings, span.End(err)
		}

		readings[url] = append(readings[url], s)
	}

	n := 0
	for _, r := range readings {
		n += len(r)
	}
	span = span.Field("stations_count", len(readings))
	span = span.Field("readings_count", n)
	return readings, span.End()
}

// 2016-01-30T00:00:00Z,http://environment.data.gov.uk/flood-monitoring/id/measures/0569TH-level-stage-i-15_min-mASD,3.430
func csvRecordToReading(r []string) (string, gauge.Reading, error) {
	var s gauge.Reading
	var err error

	if len(r) != 3 {
		return "", s, errors.New(strconv.Itoa(len(r)) + " rows in " + strings.Join(r, ","))
	}

	s.EventTime, err = time.Parse(time.RFC3339, r[0])
	if err != nil {
		return "", s, errors.New(r[0] + " is not RFC3339")
	}

	v, err := strconv.ParseFloat(r[2], 32)
	if err != nil {
		return "", s, errors.New(r[2] + " is not a float")
	}
	s.Value = float32(v)

	return r[1], s, nil
}
