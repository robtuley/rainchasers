package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func bufferAndFlushToCsv(snapC <-chan gauge.Snapshot) (<-chan string, <-chan error, error) {
	csvC := make(chan string)
	errC := make(chan error)

	n := 0
	w := csv.NewWriter(os.Stdout)
	for s := range snapC {
		n += 1
		r := []string{
			s.Url,
			s.StationUrl,
			s.Name,
			s.RiverName,
			strconv.FormatFloat(float64(s.Lat), 'f', -1, 32),
			strconv.FormatFloat(float64(s.Lg), 'f', -1, 32),
			s.Type,
			s.Unit,
			strconv.FormatInt(s.DateTime.Unix(), 10),
			strconv.FormatFloat(float64(s.Value), 'f', -1, 32),
		}
		if err := w.Write(r); err != nil {
			errC <- err
		}
		if n > 10 {
			break
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		errC <- err
	}

	return csvC, errC, nil
}
