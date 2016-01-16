package main

import (
	"encoding/json"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

type detailStationMeasureJson struct {
	Url           string          `json:"@id"`
	Name          string          `json:"label"`
	Type          string          `json:"parameter"`
	Unit          string          `json:"unit"`
	LatestRawJson json.RawMessage `json:"latestReading"`
	LatestParsed  detailStationMeasureLatestJson
}

type detailStationMeasureLatestJson struct {
	DateTime time.Time `json:"dateTime"`
	Value    float32   `json:"value"`
}

type detailStationJson struct {
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
func requestStationDetail(url string) ([]gauge.Snapshot, error) {
	waitOnApiRequestSchedule()

	defer report.Tock(report.Tick(), "detail.response", report.Data{
		"url": url,
	})

	resp, err := doJsonRequest(url)
	if err != nil {
		report.Action("detail.request.error", report.Data{"url": url, "error": err.Error()})
		return []gauge.Snapshot{}, err
	} else {
		defer resp.Body.Close()
	}

	s := detailStationJson{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&s)
	if err != nil {
		report.Action("detail.decode.error", report.Data{"url": url, "error": err.Error()})
		return []gauge.Snapshot{}, err
	}

	// the EA API returns either an ARRAY of measures OR a single OBJECT
	// in the 'measures' key which makes it hard to decode. We try decoding
	// as an array first, then failback to an single object.
	var measureArray []detailStationMeasureJson
	err = json.Unmarshal(s.Items.MeasuresRawJson, &measureArray)
	if err != nil {
		var measureObject detailStationMeasureJson
		err = json.Unmarshal(s.Items.MeasuresRawJson, &measureObject)
		measureArray = []detailStationMeasureJson{measureObject}
	}
	if err != nil {
		report.Action("detail.decode.error", report.Data{"url": url, "error": err.Error()})
		return []gauge.Snapshot{}, err
	}

	snapshots := make([]gauge.Snapshot, 0, len(measureArray))
	for _, m := range measureArray {

		// in the EA API, most latestReading keys are an object with dateTime
		// and value fields, but sometimes it doesn't -- URLs seem common. We
		// do a conditional parse and simply dump those that can't match.
		err := json.Unmarshal(m.LatestRawJson, &m.LatestParsed)
		if err != nil {
			report.Info("detail.corrupt.latestreading",
				report.Data{"url": m.Url, "station": s.Items.Url, "json": m.LatestRawJson})
			continue
		}

		v, u := normaliseUnit(m.LatestParsed.Value, m.Unit)
		if len(m.Unit) > 0 && len(u) == 0 {
			report.Action("detail.unit.error", report.Data{"url": m.Url, "unit": m.Unit})
		}

		switch m.Type {
		case "flow", "level", "temperature":
		case "wind":
			// skip wind measures as not required
			continue
		default:
			report.Action("detail.type.error", report.Data{"url": m.Url, "type": m.Type})
			continue
		}

		snapshots = append(snapshots, gauge.Snapshot{
			Url:        m.Url,
			StationUrl: s.Items.Url,
			Name:       s.Items.Name,
			RiverName:  s.Items.RiverName,
			Lat:        s.Items.Lat,
			Lg:         s.Items.Lg,
			Type:       m.Type,
			Unit:       u,
			DateTime:   m.LatestParsed.DateTime,
			Value:      v,
		})

	}

	return snapshots, nil
}

func normaliseUnit(value float32, qudtUnit string) (float32, string) {
	switch qudtUnit {
	case "http://qudt.org/1.1/vocab/unit#CubicMeterPerSecond":
		return value, "cumec"
	case "http://qudt.org/1.1/vocab/unit#DegreeCentigrade":
		return value, "centigrade"
	case "http://qudt.org/1.1/vocab/unit#Meter":
		return value, "metre"
	case "http://qudt.org/1.1/vocab/unit#MeterPerSecond":
		return value, "metre_per_second"

	case "http://qudt.org/1.1/vocab/unit#Knot":
		return value * 0.514444, "metre_per_second"
	case "http://qudt.org/1.1/vocab/unit#MegaLiterPerDay":
		return value * 0.0115741, "cumec"
	case "http://qudt.org/1.1/vocab/unit#Millimeter":
		return value * 0.001, "metre"
	case "http://qudt.org/1.1/vocab/unit#LiterPerSecond":
		return value * 0.001, "cumec"
	}

	return value, ""
}
