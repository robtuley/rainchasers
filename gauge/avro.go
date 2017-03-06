package gauge

import (
	"bytes"
	"time"

	"github.com/linkedin/goavro"
)

var avroSchemaJSON string
var avroCodec goavro.Codec

func init() {
	avroSchemaJSON = `
{ "namespace": "com.rainchasers.gauge",
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
      "type": {"type": "enum", "name": "unitValues",
                "symbols": ["", "metre", "centigrade", "cumec", "metre_per_second"]},
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

	var err error
	avroCodec, err = goavro.NewCodec(avroSchemaJSON)
	if err != nil {
		panic(err)
	}
}

func Encode(s Snapshot) (*bytes.Buffer, error) {
	bb := new(bytes.Buffer)

	// todo: is this repeated call of RecordSchema inefficient?
	r, err := goavro.NewRecord(goavro.RecordSchema(avroSchemaJSON))
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

	if err = avroCodec.Encode(bb, r); err != nil {
		return bb, err
	}

	return bb, nil
}

func Decode(bb *bytes.Buffer) (Snapshot, error) {
	var s Snapshot

	decoded, err := avroCodec.Decode(bb)
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
