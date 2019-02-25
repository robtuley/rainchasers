package request

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

const (
	httpUserAgent = "Rainchaser Bot <hello@rainchasers.com>"
)

// JSON makes a request for JSON data
func JSON(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Add("Accept", "application/json")
	req.Header.Set("User-Agent", httpUserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, errors.New("Status code " + strconv.Itoa(resp.StatusCode) + " for " + url)
	}

	return resp, nil
}

// CSV makes a request for a CSV file content
func CSV(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", httpUserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, errors.New("Status code " + strconv.Itoa(resp.StatusCode) + " for " + url)
	}

	return resp, nil
}
