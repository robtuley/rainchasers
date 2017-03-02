package main

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func getReadings(s gauge.Snapshot) ([]gauge.SnapshotUpdate, error) {
	var updates []gauge.SnapshotUpdate

	resp, err := http.Get(s.Url)
	if err != nil {
		return updates, err
	}
	if resp.StatusCode != http.StatusOK {
		return updates, errors.New(s.Url + " status code " + strconv.Itoa(resp.StatusCode))
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

		u := gauge.SnapshotUpdate{
			Url: s.Url,
		}

		updates = append(updates, u)
	}

	return updates, nil
}
