package main

import (
	"encoding/csv"
	"errors"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"io"
	"strconv"
	"strings"
	"time"
)

func getReadings(dataURL string) ([]gauge.SnapshotUpdate, error) {
	var updates []gauge.SnapshotUpdate

	resp, err := requestCSV(dataURL)
	if err != nil {
		return updates, err
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
			return updates, err
		}
		if len(r) != 2 {
			return updates, errors.New(strconv.Itoa(len(r)) + " rows in " + strings.Join(r, ","))
		}

		u := gauge.SnapshotUpdate{}

		u.DateTime, err = time.Parse("02/01/2006 15:04:05", r[0])
		if err != nil {
			continue
		}

		v, err := strconv.ParseFloat(r[1], 32)
		if err != nil {
			continue
		}
		u.Value = float32(v)

		updates = append(updates, u)
	}

	return updates, nil
}
