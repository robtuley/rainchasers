package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/report"
	"net/http"
	"strconv"
	"time"
)

func attemptBootstrap(url string, cache *gauge.Cache, log *report.Logger) {
	if len(url) == 0 {
		log.Info("bootstrap.skipped", report.Data{})
		return
	}

	tick := log.Tick()
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		<-log.Action("bootstrap.failed", report.Data{
			"url":   url,
			"error": err.Error(),
			"step":  "download.setup",
		})
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		<-log.Action("bootstrap.failed", report.Data{
			"url":   url,
			"error": err.Error(),
			"step":  "download.request",
		})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		<-log.Action("bootstrap.failed", report.Data{
			"url":   url,
			"error": "Status code " + strconv.Itoa(resp.StatusCode),
			"step":  "download.request",
		})
		return
	}
	log.Tock(tick, "bootstrap.downloaded", report.Data{
		"url": url,
		"len": resp.ContentLength,
	})

	_ = cache
}
