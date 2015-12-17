package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/robtuley/report"
)

const (
	API_REQUESTS_PER_SECOND = 1
	API_USER_AGENT          = "Rainchasers v0.1"
)

var apiRequestThrottle <-chan time.Time

func main() {
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "ea_monitor"})
	report.RuntimeStatsEvery(30 * time.Second)

	apiRequestThrottle = time.Tick(time.Second / API_REQUESTS_PER_SECOND)
	wiskiIdC := discoverEaGaugeWiskiIDs()
	n := 0
	for range wiskiIdC {
		n = n + 1
	}
	report.Info("discover.found", report.Data{"count": n})
}

func discoverEaGaugeWiskiIDs() chan string {
	idC := make(chan string)

	go func() {
		type StationList struct {
			Items []struct {
				WiskiID string `json:"wiskiId"`
			} `json:"items"`
		}

		batchSize := 100
		lastBatchSize := batchSize
		currentOffset := 0
		baseUrl := "http://environment.data.gov.uk/flood-monitoring/id/stations" +
			"?_limit=" + strconv.Itoa(batchSize)
		client := &http.Client{}

		// The paging _limit and _offset parameters apply to the number of 'measures'
		// in the EA API result set rather than the number of items, so simply iterate
		// until we receive a completely empty set.
		for lastBatchSize > 0 {
			<-apiRequestThrottle

			url := baseUrl + "&_offset=" + strconv.Itoa(currentOffset)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				report.Action("discover.request", report.Data{"error": err.Error()})
				return
			}
			req.Header.Add("Accept", "application/json")
			req.Header.Set("User-Agent", API_USER_AGENT)
			resp, err := client.Do(req)
			if err != nil {
				report.Action("discover.request", report.Data{"error": err.Error()})
				return
			}
			decoder := json.NewDecoder(resp.Body)
			s := StationList{}
			err = decoder.Decode(&s)
			resp.Body.Close()
			if err != nil {
				report.Action("discover.decode", report.Data{"error": err.Error()})
				return
			}

			for _, item := range s.Items {
				idC <- item.WiskiID
			}

			lastBatchSize = len(s.Items)
			report.Info("discover.request", report.Data{"count": lastBatchSize})
			currentOffset = currentOffset + batchSize
		}

		close(idC)
	}()

	return idC
}
