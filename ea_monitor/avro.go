package main

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
      "doc": "URL for the gauge measurement",
      "type": "string",
      "name": "url"
    },
    {
      "doc": "URL for the measuring station",
      "type": "string",
      "name": "station_url"
    },
    {
      "doc": "Human-readable name of the measurement",
      "type": "string",
      "name": "name"
    },
    {
      "doc": "Name of the river measured",
      "type": "string",
      "name": "river_name"
    },
    {
      "doc": "Location latitude",
      "type": "float",
      "name": "lat"
    },
    {
      "doc": "Location longitude",
      "type": "float",
      "name": "lg"
    },
    {
      "doc": "Measurement value",
      "type": "float",
      "name": "value"
    },
    {
      "doc": "Unix epoch time in seconds for snapshot",
      "type": "long",
      "name": "timestamp"
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

func avroEncode(s Snapshot) (*bytes.Buffer, error) {
	bb := new(bytes.Buffer)

	// todo: is this repeated call of RecordSchema inefficient?
	r, err := goavro.NewRecord(goavro.RecordSchema(avroSchemaJSON))
	if err != nil {
		return bb, err
	}

	r.Set("url", s.Url)
	r.Set("station_url", s.StationUrl)
	r.Set("name", s.Name)
	r.Set("river_name", s.RiverName)
	r.Set("lat", s.Lat)
	r.Set("lg", s.Lg)
	r.Set("value", s.Value)
	r.Set("timestamp", s.DateTime.Unix())

	if err = avroCodec.Encode(bb, r); err != nil {
		return bb, err
	}

	return bb, nil
}

func avroDecode(bb *bytes.Buffer) (Snapshot, error) {
	var s Snapshot

	decoded, err := avroCodec.Decode(bb)
	if err != nil {
		return s, err
	}

	r := decoded.(*goavro.Record)

	url, err := r.Get("url")
	if err != nil {
		return s, err
	}

	stationUrl, err := r.Get("station_url")
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

	timestamp, err := r.Get("timestamp")
	if err != nil {
		return s, err
	}

	value, err := r.Get("value")
	if err != nil {
		return s, err
	}

	s = Snapshot{
		url.(string),
		stationUrl.(string),
		name.(string),
		riverName.(string),
		lat.(float32),
		lg.(float32),
		"",
		"",
		time.Unix(timestamp.(int64), 0),
		value.(float32),
	}

	return s, nil
}
