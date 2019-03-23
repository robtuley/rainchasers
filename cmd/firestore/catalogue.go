package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/rainchasers/report"
)

// CatalogueJSON is the JSON format to download the river catalogue
// from https://app.rainchasers.com/catalogue.json
type CatalogueJSON struct {
	Version  string    `json:"version"`
	Sections []Section `json:"rivers"`
}

// Update polls for updated content version data
func parseCatalogue(ctx context.Context) ([]Section, report.Span) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url := "https://app.rainchasers.com/catalogue.json"
	span := report.StartSpan("rivers.update").Field("url", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, span.End(err)
	}
	req = req.WithContext(ctx)

	client := &http.Client{
		Timeout: time.Second * 40,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, span.End(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := errors.New("Status code " + strconv.Itoa(resp.StatusCode))
		return nil, span.End(err)
	}

	decoder := json.NewDecoder(resp.Body)
	data := CatalogueJSON{}
	if err := decoder.Decode(&data); err != nil {
		return nil, span.End(err)
	}
	span = span.Field("version", data.Version)
	span = span.Field("stations_count", len(data.Sections))

	return data.Sections, span.End()
}
