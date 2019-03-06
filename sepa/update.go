package main

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/request"
	"github.com/rainchasers/report"
)

func getReadings(d *daemon.Supervisor, dataURL string) (readings []gauge.Reading, err error) {
	ctx, cancel := context.WithTimeout(d.Context(), 60*time.Second)
	ctx = d.StartSpan(ctx, "station.http")
	defer func() {
		d.EndSpan(ctx, err, report.Data{
			"url":   dataURL,
			"count": len(readings),
		})
		cancel()
	}()

	resp, err := request.CSV(ctx, dataURL)
	if err != nil {
		return readings, err
	}
	defer resp.Body.Close()

	csv := csv.NewReader(resp.Body)

ReadCSV:
	for {
		r, err := csv.Read()

		if err == io.EOF || err == io.ErrUnexpectedEOF || err == io.ErrClosedPipe {
			break ReadCSV
		}
		if err != nil {
			return readings, err
		}
		if len(r) != 2 {
			return readings, errors.New(strconv.Itoa(len(r)) + " rows in " + strings.Join(r, ","))
		}

		u := gauge.Reading{}

		u.EventTime, err = time.Parse("02/01/2006 15:04:05", r[0])
		if err != nil {
			continue
		}

		v, err := strconv.ParseFloat(r[1], 32)
		if err != nil {
			continue
		}
		u.Value = float32(v)

		readings = append(readings, u)
	}

	return readings, nil
}
