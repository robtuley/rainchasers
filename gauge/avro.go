package gauge

import (
	"bytes"
	"fmt"
	"github.com/linkedin/goavro"
	"time"
)

const ReadingSchema = `
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
	SnapshotSchema string
	snapshotCodec  goavro.Codec
)

func init() {
	SnapshotSchema = fmt.Sprintf(`
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
`, ReadingSchema)

	var err error

	snapshotCodec, err = goavro.NewCodec(SnapshotSchema)
	if err != nil {
		panic(err)
	}
}

func (s *Snapshot) Encode() (*bytes.Buffer, error) {
	bb := new(bytes.Buffer)
	var innerRecords []interface{}

	for _, u := range s.Readings {
		m, err := goavro.NewRecord(goavro.RecordSchema(ReadingSchema))
		if err != nil {
			return bb, err
		}

		m.Set("timestamp", u.DateTime.Unix())
		m.Set("value", u.Value)
		innerRecords = append(innerRecords, m)
	}

	outerRecord, err := goavro.NewRecord(goavro.RecordSchema(SnapshotSchema))
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

func (s *Snapshot) Decode(bb *bytes.Buffer) error {
	decoded, err := snapshotCodec.Decode(bb)
	if err != nil {
		return err
	}

	r := decoded.(*goavro.Record)

	dataURL, err := r.Get("data_url")
	if err != nil {
		return err
	}

	humanURL, err := r.Get("human_url")
	if err != nil {
		return err
	}

	name, err := r.Get("name")
	if err != nil {
		return err
	}

	riverName, err := r.Get("river_name")
	if err != nil {
		return err
	}

	lat, err := r.Get("lat")
	if err != nil {
		return err
	}

	lg, err := r.Get("lg")
	if err != nil {
		return err
	}

	typeEnum, err := r.Get("type")
	if err != nil {
		return err
	}

	unit, err := r.Get("unit")
	if err != nil {
		return err
	}

	s.Station = Station{
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
		return err
	}
	innerRecords := data.([]interface{})
	for _, a := range innerRecords {
		u := a.(*goavro.Record)

		timestamp, err := u.Get("timestamp")
		if err != nil {
			return err
		}

		value, err := u.Get("value")
		if err != nil {
			return err
		}

		s.Readings = append(s.Readings, Reading{
			DateTime: time.Unix(timestamp.(int64), 0),
			Value:    value.(float32),
		})
	}

	return nil
}
