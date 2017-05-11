package main

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func bootstrapSnapshots(url string) (*bytes.Buffer, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return new(bytes.Buffer), err
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return new(bytes.Buffer), err
	}
	if resp.StatusCode != http.StatusOK {
		return new(bytes.Buffer), errors.New("Status code " + strconv.Itoa(resp.StatusCode) + " for " + url)
	}

	bb := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, err = bb.ReadFrom(resp.Body)

	return bb, nil
}
