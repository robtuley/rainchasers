package main

import (
	"encoding/json"

	"github.com/robtuley/report"
)

// Retrieve the detail and latest readings for an individual gauge.
func requestStationDetail(url string, reTryC chan string, resultC chan GaugeSnapshot) {
	waitOnApiRequestSchedule()

	go func() {
		type StationIndividual struct {
			Items struct {
				Url       string  `json:"@id"`
				Name      string  `json:"label"`
				RiverName string  `json:"riverName"`
				Lat       float32 `json:"lat"`
				Lg        float32 `json:"long"`
			} `json:"items"`
		}

		tick := report.Tick()

		err, resp := doJsonRequest(url)
		if err != nil {
			report.Action("detail.request.error", report.Data{"url": url, "error": err.Error()})
			reTryC <- url
			return
		} else {
			defer resp.Body.Close()
		}

		s := StationIndividual{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&s)
		if err != nil {
			report.Action("detail.decode.error", report.Data{"url": url, "error": err.Error()})
			reTryC <- url
			return
		}

		report.Tock(tick, "detail.response", report.Data{
			"url": url,
		})

		resultC <- GaugeSnapshot{
			s.Items.Url, s.Items.Name, s.Items.RiverName, s.Items.Lat, s.Items.Lg,
		}
	}()
}
