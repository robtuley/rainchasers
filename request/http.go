package request

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

const (
	httpTimeoutInSeconds = 60
	httpUserAgent        = "Rainchaser Bot <hello@rainchasers.com>"
)

func JSON(url string) (*http.Response, error) {
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

func CSV(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: httpTimeoutInSeconds * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", httpUserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, errors.New("Status code " + strconv.Itoa(resp.StatusCode) + " for " + url)
	}

	return resp, nil
}
