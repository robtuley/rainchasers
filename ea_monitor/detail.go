package main

import (
	"encoding/json"

	"github.com/robtuley/report"
)

// Retrieve the detail and latest readings for an individual gauge.
func requestStationDetail(wiskiId string, reTryC chan string, resultC chan GaugeSnapshot) {
	waitOnApiRequestSchedule()

	go func() {
		type StationIndividual struct {
			Items struct {
				WiskiID   string  `json:"wiskiId"`
				Name      string  `json:"label"`
				RiverName string  `json:"riverName"`
				Lat       float32 `json:"lat"`
				Lg        float32 `json:"long"`
			} `json:"items"`
		}

		url := "http://environment.data.gov.uk/flood-monitoring/id/stations/" + wiskiId
		defer report.Tock(report.Tick(), "detail.response", report.Data{
			"url": url,
		})

		err, resp := doJsonRequest(url)
		if err != nil {
			report.Action("detail.request.error", report.Data{"url": url, "error": err.Error()})
			reTryC <- wiskiId
			return
		} else {
			defer resp.Body.Close()
		}

		s := StationIndividual{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&s)
		if err != nil {
			report.Action("detail.decode.error", report.Data{"url": url, "error": err.Error()})
			reTryC <- wiskiId
			return
		}

		resultC <- GaugeSnapshot{
			s.Items.WiskiID, s.Items.Name, s.Items.RiverName, s.Items.Lat, s.Items.Lg,
		}

	}()
}
