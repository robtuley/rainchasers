package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/report"
)

func bootstrapGaugeCache(url string, cache *gauge.Cache, log *report.Logger) <-chan struct{} {
	doneC := make(chan struct{})

	if len(url) == 0 {
		log.Info("bootstrap.gauge.skipped", report.Data{})
		close(doneC)
		return doneC
	}

	go func() {
		defer close(doneC)

		tick := log.Tick()
		client := &http.Client{
			Timeout: 15 * time.Second,
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			<-log.Action("bootstrap.gauge.failed", report.Data{
				"url":   url,
				"error": err.Error(),
				"step":  "setup",
			})
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			<-log.Action("bootstrap.gauge.failed", report.Data{
				"url":   url,
				"error": err.Error(),
				"step":  "request",
			})
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			<-log.Action("bootstrap.gauge.failed", report.Data{
				"url":   url,
				"error": "Status code " + strconv.Itoa(resp.StatusCode),
				"step":  "request",
			})
			return
		}
		log.Tock(tick, "bootstrap.gauge.downloaded", report.Data{
			"url": url,
			"len": resp.ContentLength,
		})

		if err := cache.Decode(resp.Body); err != nil {
			<-log.Action("bootstrap.gauge.failed", report.Data{
				"url":   url,
				"error": err.Error(),
				"step":  "decode",
			})
			return
		}
		log.Tock(tick, "bootstrap.gauge.complete", report.Data{"url": url})
	}()

	return doneC
}
