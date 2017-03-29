package queue

import (
	"bytes"
	"fmt"
	"github.com/linkedin/goavro"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"time"
)

const readingSchemaJSON = `
{
  "namespace": "com.rainchasers.gauge",
  "type": "record",
  "name": "measure",
  "doc:": "Gauge measurement information",
  "fields": [
    {
      "doc": "Unix epoch time in seconds for measurement",
      "type": "long",
      "name": "timestamp"
    },{
      "doc": "Measurement value",
      "type": "float",
      "name": "value"
    }
  ]
}
`

var (
	snapshotSchemaJSON string
	snapshotCodec      goavro.Codec
)

func init() {
	snapshotSchemaJSON = fmt.Sprintf(`
{
  "namespace": "com.rainchasers.gauge",
  "type": "record",
  "name": "snapshot",
  "doc:": "Gauge measurement record information and reading snapshot",
  "fields": [
    {
      "doc": "Data URL for the gauge measurement",
      "type": "string",
      "name": "data_url"
    },{
      "doc": "Human linkable URL for the station",
      "type": "string",
      "name": "human_url"
    },{
      "doc": "Human-readable name of the measurement",
      "type": "string",
      "name": "name"
    },{
      "doc": "Name of the river measured",
      "type": "string",
      "name": "river_name"
    },{
      "doc": "Location latitude",
      "type": "float",
      "name": "lat"
    },{
      "doc": "Location longitude",
      "type": "float",
      "name": "lg"
    },{
      "doc": "Measurement unit",
      "type": "string",
      "name": "unit"
    },{
      "doc": "Measurement type",
      "type": {"type": "enum", "name": "typeValues",
                "symbols": ["level", "flow", "temperature", "rainfall"]},
      "name": "type"
    },{
      "type": {
        "items": %s,
        "type": "array"
      },
      "name": "data"
    }
  ]
}
`, readingSchemaJSON)

	var err error

	snapshotCodec, err = goavro.NewCodec(snapshotSchemaJSON)
	if err != nil {
		panic(err)
	}
}

func Encode(s *gauge.Snapshot) (*bytes.Buffer, error) {
	bb := new(bytes.Buffer)
	var innerRecords []interface{}

	for _, u := range s.Readings {
		m, err := goavro.NewRecord(goavro.RecordSchema(readingSchemaJSON))
		if err != nil {
			return bb, err
		}

		m.Set("timestamp", u.DateTime.Unix())
		m.Set("value", u.Value)
		innerRecords = append(innerRecords, m)
	}

	outerRecord, err := goavro.NewRecord(goavro.RecordSchema(snapshotSchemaJSON))
	if err != nil {
		return bb, nil
	}
	outerRecord.Set("data_url", s.Station.DataURL)
	outerRecord.Set("human_url", s.Station.HumanURL)
	outerRecord.Set("name", s.Station.Name)
	outerRecord.Set("river_name", s.Station.RiverName)
	outerRecord.Set("lat", s.Station.Lat)
	outerRecord.Set("lg", s.Station.Lg)
	outerRecord.Set("type", goavro.Enum{
		Name:  "typeValues",
		Value: s.Station.Type,
	})
	outerRecord.Set("unit", s.Station.Unit)
	outerRecord.Set("data", innerRecords)
	if err = snapshotCodec.Encode(bb, outerRecord); err != nil {
		return bb, err
	}

	return bb, nil
}

func Decode(bb *bytes.Buffer) (*gauge.Snapshot, error) {
	var s gauge.Snapshot

	decoded, err := snapshotCodec.Decode(bb)
	if err != nil {
		return &s, err
	}

	r := decoded.(*goavro.Record)

	dataURL, err := r.Get("data_url")
	if err != nil {
		return &s, err
	}

	humanURL, err := r.Get("human_url")
	if err != nil {
		return &s, err
	}

	name, err := r.Get("name")
	if err != nil {
		return &s, err
	}

	riverName, err := r.Get("river_name")
	if err != nil {
		return &s, err
	}

	lat, err := r.Get("lat")
	if err != nil {
		return &s, err
	}

	lg, err := r.Get("lg")
	if err != nil {
		return &s, err
	}

	typeEnum, err := r.Get("type")
	if err != nil {
		return &s, err
	}

	unit, err := r.Get("unit")
	if err != nil {
		return &s, err
	}

	s.Station = gauge.Station{
		DataURL:   dataURL.(string),
		HumanURL:  humanURL.(string),
		Name:      name.(string),
		RiverName: riverName.(string),
		Lat:       lat.(float32),
		Lg:        lg.(float32),
		Type:      typeEnum.(goavro.Enum).Value,
		Unit:      unit.(string),
	}

	data, err := r.Get("data")
	if err != nil {
		return &s, err
	}
	innerRecords := data.([]interface{})
	for _, a := range innerRecords {
		u := a.(*goavro.Record)

		timestamp, err := u.Get("timestamp")
		if err != nil {
			return &s, err
		}

		value, err := u.Get("value")
		if err != nil {
			return &s, err
		}

		s.Readings = append(s.Readings, gauge.Reading{
			DateTime: time.Unix(timestamp.(int64), 0),
			Value:    value.(float32),
		})
	}

	return &s, nil
}
