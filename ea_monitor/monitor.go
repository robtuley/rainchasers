package main

import (
	"encoding/json"
	"net/http"
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
	for id := range wiskiIdC {
		report.Info("discover.found", report.Data{"wiskiId": id})
	}
}

func discoverEaGaugeWiskiIDs() chan string {
	idC := make(chan string)

	go func() {
		type StationList struct {
			Meta struct {
				Limit  int `json:"limit"`
				Offset int `json:"offset"`
			} `json:"meta"`
			Items []struct {
				WiskiID string `json:"wiskiId"`
			} `json:"items"`
		}

		url := "http://environment.data.gov.uk/flood-monitoring/id/stations?_limit=50"
		client := &http.Client{}

		<-apiRequestThrottle
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			report.Action("discover.request.error", report.Data{"error": err.Error()})
			return
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Set("User-Agent", API_USER_AGENT)
		resp, err := client.Do(req)
		if err != nil {
			report.Action("discover.request.error", report.Data{"error": err.Error()})
			return
		}
		report.Info("discover.dump", report.Data{"response": resp})
		decoder := json.NewDecoder(resp.Body)
		s := StationList{}
		err = decoder.Decode(&s)
		resp.Body.Close()
		if err != nil {
			report.Action("discover.decode.error", report.Data{"error": err.Error()})
			return
		}

		for _, item := range s.Items {
			idC <- item.WiskiID
		}

		close(idC)
	}()

	return idC
}
