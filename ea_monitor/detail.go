package main

import (
	"encoding/json"
	"time"

	"github.com/robtuley/report"
)

type StationMeasure struct {
	Url    string `json:"@id"`
	Name   string `json:"label"`
	Type   string `json:"parameter"`
	Unit   string `json:"unit"`
	Latest struct {
		DateTime time.Time `json:"dateTime"`
		Value    float32   `json:"value"`
	} `json:"latestReading"`
}

type StationIndividual struct {
	Items struct {
		Url             string          `json:"@id"`
		Name            string          `json:"label"`
		RiverName       string          `json:"riverName"`
		Lat             float32         `json:"lat"`
		Lg              float32         `json:"long"`
		MeasuresRawJson json.RawMessage `json:"measures"`
	} `json:"items"`
}

// Retrieve the detail and latest readings for an individual gauge.
func requestStationDetail(url string) (error, []Snapshot) {
	waitOnApiRequestSchedule()

	defer report.Tock(report.Tick(), "detail.response", report.Data{
		"url": url,
	})

	err, resp := doJsonRequest(url)
	if err != nil {
		report.Action("detail.request.error", report.Data{"url": url, "error": err.Error()})
		return err, []Snapshot{}
	} else {
		defer resp.Body.Close()
	}

	s := StationIndividual{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&s)
	if err != nil {
		report.Action("detail.decode.error", report.Data{"url": url, "error": err.Error()})
		return err, []Snapshot{}
	}

	// the EA API returns either an ARRAY of measures OR a single OBJECT
	// which makes it hard to decode. We try decoding as an array first,
	// then failback to an single object.
	var measureArray []StationMeasure
	err = json.Unmarshal(s.Items.MeasuresRawJson, &measureArray)
	if err != nil {
		var measureObject StationMeasure
		err = json.Unmarshal(s.Items.MeasuresRawJson, &measureObject)
		measureArray = []StationMeasure{measureObject}
	}
	if err != nil {
		report.Action("detail.decode.error", report.Data{"url": url, "error": err.Error()})
		return err, []Snapshot{}
	}

	snapshots := make([]Snapshot, len(measureArray))
	for k, m := range measureArray {
		snapshots[k] = Snapshot{
			m.Url,
			s.Items.Url,
			s.Items.Name,
			s.Items.RiverName,
			s.Items.Lat,
			s.Items.Lg,
			m.Type,
			m.Unit,
			m.Latest.DateTime,
			m.Latest.Value,
		}
	}

	report.Info("debug", report.Data{"measures": snapshots})

	return nil, snapshots
}
