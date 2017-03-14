package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

func doJsonRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: httpTimeoutInSeconds * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("User-Agent", httpUserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, errors.New("Status code " + strconv.Itoa(resp.StatusCode))
	}

	return resp, nil
}
