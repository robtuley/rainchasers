package main

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/rainchasers/content/internal/daemon"
	"github.com/rainchasers/content/internal/gauge"
	"github.com/rainchasers/report"
)

func getReadings(ctx context.Context, dataURL string) ([]gauge.Reading, report.Span) {
	span := report.StartSpan("sepa.recent")
	span = span.Field("url", dataURL)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	resp, err := daemon.CSV(ctx, dataURL)
	if err != nil {
		return nil, span.End(err)
	}
	defer resp.Body.Close()

	var readings []gauge.Reading
	csv := csv.NewReader(resp.Body)

ReadCSV:
	for {
		r, err := csv.Read()

		if err == io.EOF || err == io.ErrUnexpectedEOF || err == io.ErrClosedPipe {
			break ReadCSV
		}
		if err != nil {
			return readings, span.End(err)
		}
		if len(r) != 2 {
			err = errors.New(strconv.Itoa(len(r)) + " rows in " + strings.Join(r, ","))
			return readings, span.End(err)
		}

		u := gauge.Reading{}

		u.EventTime, err = time.Parse("02/01/2006 15:04:05", r[0])
		if err != nil {
			continue ReadCSV
		}

		v, err := strconv.ParseFloat(r[1], 32)
		if err != nil {
			continue ReadCSV
		}
		u.Value = float32(v)

		readings = append(readings, u)
	}

	span = span.Field("readings_count", len(readings))
	return readings, span.End()
}
