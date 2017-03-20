package gauge

import (
	"bytes"
	"fmt"
	"github.com/linkedin/goavro"
	"time"
)

const snapshotSchemaJSON = `
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
      "doc": "Unix epoch time in seconds for snapshot",
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

const measureSchemaJSON = `
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
	measurementsSchemaJSON string
	snapshotCodec          goavro.Codec
	measurementsCodec      goavro.Codec
)

func init() {
	measurementsSchemaJSON = fmt.Sprintf(`
{
  "namespace": "com.rainchasers.gauge",
  "type": "record",
  "name": "measurements",
  "doc:": "Array of gauge measurement information",
  "fields": [
    {
      "doc": "Data URL for the gauge measurement",
      "type": "string",
      "name": "data_url"
    },{
      "type": {
        "items": %s,
        "type": "array"
      },
      "name": "data"
    }
  ]
}
`, measureSchemaJSON)

	var err error

	snapshotCodec, err = goavro.NewCodec(snapshotSchemaJSON)
	if err != nil {
		panic(err)
	}

	measurementsCodec, err = goavro.NewCodec(measurementsSchemaJSON)
	if err != nil {
		panic(err)
	}
}

func EncodeSnapshot(s Snapshot) (*bytes.Buffer, error) {
	bb := new(bytes.Buffer)

	r, err := goavro.NewRecord(goavro.RecordSchema(snapshotSchemaJSON))
	if err != nil {
		return bb, err
	}

	r.Set("data_url", s.DataURL)
	r.Set("human_url", s.HumanURL)
	r.Set("name", s.Name)
	r.Set("river_name", s.RiverName)
	r.Set("lat", s.Lat)
	r.Set("lg", s.Lg)
	r.Set("type", s.Type)
	r.Set("unit", s.Unit)
	r.Set("timestamp", s.DateTime.Unix())
	r.Set("value", s.Value)

	if err = snapshotCodec.Encode(bb, r); err != nil {
		return bb, err
	}

	return bb, nil
}

func DecodeSnapshot(bb *bytes.Buffer) (Snapshot, error) {
	var s Snapshot

	decoded, err := snapshotCodec.Decode(bb)
	if err != nil {
		return s, err
	}

	r := decoded.(*goavro.Record)

	dataURL, err := r.Get("data_url")
	if err != nil {
		return s, err
	}

	humanURL, err := r.Get("human_url")
	if err != nil {
		return s, err
	}

	name, err := r.Get("name")
	if err != nil {
		return s, err
	}

	riverName, err := r.Get("river_name")
	if err != nil {
		return s, err
	}

	lat, err := r.Get("lat")
	if err != nil {
		return s, err
	}

	lg, err := r.Get("lg")
	if err != nil {
		return s, err
	}

	typeStr, err := r.Get("type")
	if err != nil {
		return s, err
	}

	unit, err := r.Get("unit")
	if err != nil {
		return s, err
	}

	timestamp, err := r.Get("timestamp")
	if err != nil {
		return s, err
	}

	value, err := r.Get("value")
	if err != nil {
		return s, err
	}

	s = Snapshot{
		DataURL:   dataURL.(string),
		HumanURL:  humanURL.(string),
		Name:      name.(string),
		RiverName: riverName.(string),
		Lat:       lat.(float32),
		Lg:        lg.(float32),
		Type:      typeStr.(string),
		Unit:      unit.(string),
		DateTime:  time.Unix(timestamp.(int64), 0),
		Value:     value.(float32),
	}

	return s, nil
}

func EncodeSnapshotUpdates(dataURL string, updates []SnapshotUpdate) (*bytes.Buffer, error) {
	bb := new(bytes.Buffer)
	var innerRecords []interface{}

	for _, u := range updates {
		m, err := goavro.NewRecord(goavro.RecordSchema(measureSchemaJSON))
		if err != nil {
			return bb, err
		}

		m.Set("timestamp", u.DateTime.Unix())
		m.Set("value", u.Value)
		innerRecords = append(innerRecords, m)
	}

	outerRecord, err := goavro.NewRecord(goavro.RecordSchema(measurementsSchemaJSON))
	if err != nil {
		return bb, nil
	}
	outerRecord.Set("data_url", dataURL)
	outerRecord.Set("data", innerRecords)
	if err = measurementsCodec.Encode(bb, outerRecord); err != nil {
		return bb, err
	}

	return bb, nil
}

func DecodeSnapshotUpdates(bb *bytes.Buffer) (string, []SnapshotUpdate, error) {
	var updates []SnapshotUpdate
	var dataURL string

	decoded, err := measurementsCodec.Decode(bb)
	if err != nil {
		return dataURL, updates, err
	}

	outerRecord := decoded.(*goavro.Record)
	url, err := outerRecord.Get("data_url")
	if err != nil {
		return dataURL, updates, err
	}
	dataURL = url.(string)

	r, err := outerRecord.Get("data")
	if err != nil {
		return dataURL, updates, err
	}
	innerRecords := r.([]interface{})
	for _, r = range innerRecords {
		u := r.(*goavro.Record)

		timestamp, err := u.Get("timestamp")
		if err != nil {
			return dataURL, updates, err
		}

		value, err := u.Get("value")
		if err != nil {
			return dataURL, updates, err
		}

		updates = append(updates, SnapshotUpdate{
			DateTime: time.Unix(timestamp.(int64), 0),
			Value:    value.(float32),
		})
	}

	return dataURL, updates, nil
}
