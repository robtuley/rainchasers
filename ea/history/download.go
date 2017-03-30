package main

import (
	"encoding/csv"
	"errors"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/util"
	"io"
	"strconv"
	"strings"
	"time"
)

func download(day time.Time) (map[string][]gauge.Reading, error) {
	url := "http://environment.data.gov.uk/flood-monitoring/archive/readings-" + day.Format("2006-01-02") + ".csv"
	snapshots := make(map[string][]gauge.Reading)

	resp, err := util.RequestCSV(url)
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

	s.DateTime, err = time.Parse(time.RFC3339, r[0])
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
