package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

const (
	API_REQUESTS_PER_SECOND = 1
	API_USER_AGENT          = "Rainchasers v0.1"
)

var apiRequestThrottle <-chan time.Time

func init() {
	apiRequestThrottle = time.Tick(time.Second / API_REQUESTS_PER_SECOND)
}

func doJsonRequest(url string) (error, *http.Response) {
	<-apiRequestThrottle
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("User-Agent", API_USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Status code " + strconv.Itoa(resp.StatusCode)), nil
	}

	return nil, resp
}
