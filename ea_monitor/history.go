package main

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func downloadHistoricalDataForDaysAgo(nDays int, updateC chan<- gauge.SnapshotUpdate) <-chan error {
	errC := make(chan error)

	go func() {
		day := time.Now().AddDate(0, 0, -1*nDays)
		url := "http://environment.data.gov.uk/flood-monitoring/archive/readings-" + day.Format("2006-01-02") + ".csv"

		resp, err := http.Get(url)
		if err != nil {
			errC <- err
			return
		}
		if resp.StatusCode != http.StatusOK {
			errC <- errors.New("Status code " + strconv.Itoa(resp.StatusCode))
			return
		}
		defer resp.Body.Close()

		csv := csv.NewReader(resp.Body)
		isFirst := true
		for {
			r, err := csv.Read()
			if err == io.EOF {
				break
			}
			// some corrupt reading values appear as 1.23|4.56 so
			// we simply skip these as known errors.
			if strings.Contains(r[2], "|") {
				continue
			}
			if err != nil {
				errC <- err
				continue
			}
			if isFirst {
				isFirst = false
				continue
			}

			s, err := csvRecordToSnapshotUpdate(r)
			if err != nil {
				errC <- err
				continue
			}

			updateC <- s
		}

		close(errC)
	}()

	return errC
}

// 2016-01-30T00:00:00Z,http://environment.data.gov.uk/flood-monitoring/id/measures/0569TH-level-stage-i-15_min-mASD,3.430
func csvRecordToSnapshotUpdate(r []string) (gauge.SnapshotUpdate, error) {
	var s gauge.SnapshotUpdate
	var err error

	if len(r) != 3 {
		return s, errors.New(strconv.Itoa(len(r)) + " rows in " + strings.Join(r, ","))
	}
	s.Url = r[1]

	s.DateTime, err = time.Parse(time.RFC3339, r[0])
	if err != nil {
		return s, errors.New(r[0] + " is not RFC3339")
	}

	v, err := strconv.ParseFloat(r[2], 32)
	if err != nil {
		return s, errors.New(r[2] + " is not a float")
	}
	s.Value = float32(v)

	return s, nil
}
