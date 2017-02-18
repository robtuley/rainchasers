package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
		Url              string `json:"@id"`
		Name             string
		NameRawJson      json.RawMessage `json:"label"`
		RiverName        string
		RiverNameRawJson json.RawMessage `json:"riverName"`
		Lat              float32
		Lg               float32
		LatRawJson       json.RawMessage `json:"lat"`
		LgRawJson        json.RawMessage `json:"long"`
		MeasuresRawJson  json.RawMessage `json:"measures"`
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
		if resp != nil {
			if resp.StatusCode == http.StatusNotFound {
				report.Info("detail.request.notfound", report.Data{"url": url, "error": err.Error()})
				return []gauge.Snapshot{}, err
			}
		}
		report.Action("detail.request.error", report.Data{"url": url, "error": err.Error()})
		return []gauge.Snapshot{}, err
	} else {
		defer resp.Body.Close()
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		report.Action("detail.body.error", report.Data{"url": url, "error": err.Error()})
		return []gauge.Snapshot{}, err
	}

	s := detailStationJson{}
	err = json.Unmarshal(bodyBytes, &s)
	if err != nil {
		report.Action("detail.parse.error", report.Data{"url": url, "error": err.Error()})
		return []gauge.Snapshot{}, err
	}
	// a known inconsistency is that the API can provide Lat, Lg or label as an array
	// so we use a defensive mechanism to parse these fields and let them be missing completely
	// e.g. http://environment.data.gov.uk/flood-monitoring/id/stations/E40411
	s.Items.Lat, err = parseFloatFromScalarOrArray(s.Items.LatRawJson)
	if err != nil {
		report.Info("detail.lat.missing", report.Data{"url": url, "error": err.Error()})
		s.Items.Lat = 0.0
	}
	s.Items.Lg, err = parseFloatFromScalarOrArray(s.Items.LgRawJson)
	if err != nil {
		report.Info("detail.lg.missing", report.Data{"url": url, "error": err.Error()})
		s.Items.Lg = 0.0
	}
	s.Items.Name, err = parseStringFromScalarOrArray(s.Items.NameRawJson)
	if err != nil {
		report.Info("detail.name.missing", report.Data{"url": url, "error": err.Error()})
		s.Items.Name = ""
	}
	s.Items.RiverName, err = parseStringFromScalarOrArray(s.Items.RiverNameRawJson)
	if err != nil {
		report.Info("detail.rivername.missing", report.Data{"url": url, "error": err.Error()})
		s.Items.RiverName = ""
	}

	// the EA API returns either an:
	//   * missing 'measures' key
	//   * ARRAY in the measures key
	//   * single OBJECT in the measures key
	// This makes decoding complex.
	if s.Items.MeasuresRawJson == nil {
		report.Info("detail.measure.missing", report.Data{"url": url})
		return []gauge.Snapshot{}, nil
	}
	var measureArray []detailStationMeasureJson
	err = json.Unmarshal(s.Items.MeasuresRawJson, &measureArray)
	if err != nil {
		var measureObject detailStationMeasureJson
		err = json.Unmarshal(s.Items.MeasuresRawJson, &measureObject)
		measureArray = []detailStationMeasureJson{measureObject}
	}
	if err != nil {
		report.Action("detail.measure.error", report.Data{"url": url, "error": err.Error()})
		return []gauge.Snapshot{}, err
	}

	snapshots := make([]gauge.Snapshot, 0, len(measureArray))
	for _, m := range measureArray {

		// in the EA API, most latestReading keys are an object with dateTime
		// and value fields, but sometimes it isn't -- URLs seem common. We
		// do a conditional parse and simply dump those that can't match.
		err := json.Unmarshal(m.LatestRawJson, &m.LatestParsed)
		if err != nil {
			report.Info("detail.latestreading.corrupt",
				report.Data{"url": m.Url, "station": s.Items.Url, "json": string(m.LatestRawJson)})
			continue
		}

		v, u := normaliseUnit(m.LatestParsed.Value, m.Unit)
		if len(m.Unit) > 0 && len(u) == 0 {
			report.Action("detail.unit.error", report.Data{"url": m.Url, "unit": m.Unit})
		}

		switch m.Type {
		case "flow", "level", "temperature", "rainfall":
		case "wind":
			// skip wind measures as not required
			continue
		default:
			report.Action("detail.type.error", report.Data{"url": m.Url, "unknownType": m.Type})
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
