package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

const (
	API_REQUESTS_PER_SECOND = 1
	API_USER_AGENT          = "Rainchaser Bot <hello@rainchasers.com>"
)

var apiRequestThrottle <-chan time.Time

func init() {
	apiRequestThrottle = time.Tick(time.Second / API_REQUESTS_PER_SECOND)
}

func waitOnApiRequestSchedule() {
	<-apiRequestThrottle
}

func doJsonRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("User-Agent", API_USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Status code " + strconv.Itoa(resp.StatusCode))
	}

	return resp, nil
}
