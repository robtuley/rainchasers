package ea

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/report"
)

// Day downloads all the measurements on specified day
func Day(d *daemon.Supervisor, day time.Time) (rd map[string][]gauge.Reading, err error) {
	url := "http://environment.data.gov.uk/flood-monitoring/archive/readings-" + day.Format("2006-01-02") + ".csv"
	snapshots := make(map[string][]gauge.Reading)

	ctx, cancel := context.WithTimeout(d.Context(), 60*time.Second)
	ctx = d.StartSpan(ctx, "ea.day")
	defer func() {
		n := 0
		for _, r := range rd {
			n += len(r)
		}
		d.EndSpan(ctx, err, report.Data{
			"count_stations": len(rd),
			"count_readings": n,
			"url":            url,
		})
		cancel()
	}()

	resp, err := daemon.CSV(ctx, url)
	if err != nil {
		return snapshots, err
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
			return snapshots, err
		}
		if isFirst {
			isFirst = false
			continue
		}

		url, s, err := csvRecordToReading(r)
		if err != nil {
			return snapshots, err
		}

		snapshots[url] = append(snapshots[url], s)
	}

	return snapshots, nil
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
