package main

import (
	"encoding/json"
	"strconv"

	"github.com/robtuley/report"
)

// Discover all available EA identifying URLs.
//
//     for url := range discoverUrls() {
//         log.Println(url)
//     }
//
func discoverStationUrls() chan string {
	urlC := make(chan string)

	go func() {
		type StationList struct {
			Items []struct {
				Url string `json:"@id"`
			} `json:"items"`
		}

		batchSize := 100
		lastBatchSize := batchSize
		currentOffset := 0
		baseUrl := "http://environment.data.gov.uk/flood-monitoring/id/stations" +
			"?_limit=" + strconv.Itoa(batchSize)

		// The paging _limit and _offset parameters apply to the number of 'measures'
		// in the EA API result set rather than the number of items, so simply iterate
		// until we receive a completely empty set.
		for lastBatchSize > 0 {
			waitOnApiRequestSchedule()

			url := baseUrl + "&_offset=" + strconv.Itoa(currentOffset)
			currentOffset = currentOffset + batchSize
			tick := report.Tick()

			err, resp := doJsonRequest(url)
			if err != nil {
				report.Action("discover.request.error", report.Data{"url": url, "error": err.Error()})
				resp.Body.Close()
				continue
			}

			s := StationList{}
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&s)
			resp.Body.Close()
			if err != nil {
				report.Action("discover.decode.error", report.Data{"url": url, "error": err.Error()})
				continue
			}

			lastBatchSize = len(s.Items)
			report.Tock(tick, "discovery.response", report.Data{
				"count": lastBatchSize,
				"url":   url,
			})

			for _, item := range s.Items {
				urlC <- item.Url
			}
		}

		close(urlC)
	}()

	return urlC
}
