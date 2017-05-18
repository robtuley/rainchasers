package gauge

import (
	"fmt"
	"io"
	"time"

	"github.com/linkedin/goavro"
)

// ReadingSchema is the AVRO schema for a Reading
const ReadingSchema = `
{
  "namespace": "com.rainchasers.gauge",
  "type": "record",
  "name": "measure",
  "doc:": "Gauge measurement information",
  "fields": [
    {
      "doc": "Unix epoch time in seconds for measurement event time",
      "type": "long",
      "name": "event_time"
    },{
      "doc": "Measurement value",
      "type": "float",
      "name": "value"
    }
  ]
}
`

var (
	// SnapshotSchema is the AVRO schema for a Snapshot
	SnapshotSchema string
	snapshotCodec  goavro.Codec
	// CacheSchema is the AVRO schema for a Cache
	CacheSchema string
	cacheCodec  goavro.Codec
)

func init() {
	var err error

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
      "name": "readings"
    },{
      "doc": "Unix epoch time in seconds for measurement processing time",
      "type": "long",
      "name": "processing_time"
    }
  ]
}
`, ReadingSchema)

	snapshotCodec, err = goavro.NewCodec(SnapshotSchema)
	if err != nil {
		panic(err)
	}

	CacheSchema = fmt.Sprintf(`
{
  "namespace": "com.rainchasers.gauge",
  "type": "record",
  "name": "cache",
  "doc:": "Gauge measurement station and reading cache",
  "fields": [
    {
      "type": {
        "items": %s,
        "type": "array"
      },
      "name": "snapshots"
    }
  ]
}
`, SnapshotSchema)

	cacheCodec, err = goavro.NewCodec(CacheSchema)
	if err != nil {
		panic(err)
	}
}

func readingToRecord(r *Reading) (*goavro.Record, error) {
	record, err := goavro.NewRecord(goavro.RecordSchema(ReadingSchema))
	if err != nil {
		return record, err
	}
	record.Set("event_time", r.EventTime.Unix())
	record.Set("value", r.Value)
	return record, nil
}

func snapshotToRecord(s *Snapshot) (*goavro.Record, error) {
	outerRecord, err := goavro.NewRecord(goavro.RecordSchema(SnapshotSchema))
	if err != nil {
		return outerRecord, err
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

	var innerRecords []interface{}
	for i := range s.Readings {
		m, err := readingToRecord(&s.Readings[i])
		if err != nil {
			return outerRecord, err
		}
		innerRecords = append(innerRecords, m)
	}
	outerRecord.Set("readings", innerRecords)
	outerRecord.Set("processing_time", s.ProcessingTime.Unix())

	return outerRecord, nil
}

// Encode writes out the Snapshot in AVRO binary format
func (s *Snapshot) Encode(writer io.Writer) error {
	record, err := snapshotToRecord(s)
	if err != nil {
		return err
	}

	if err := snapshotCodec.Encode(writer, record); err != nil {
		return err
	}

	return nil
}

// Encode writes out the Cache in AVRO binary format
func (c *Cache) Encode(writer io.Writer) error {
	outerRecord, err := goavro.NewRecord(goavro.RecordSchema(CacheSchema))
	if err != nil {
		return err
	}

	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	var innerRecords []interface{}
	for i := range c.snapMap {
		s, err := snapshotToRecord(c.snapMap[i])
		if err != nil {
			return err
		}
		innerRecords = append(innerRecords, s)
	}
	outerRecord.Set("snapshots", innerRecords)

	if err = cacheCodec.Encode(writer, outerRecord); err != nil {
		return err
	}

	return nil
}

func recordToStation(r *goavro.Record) (Station, error) {
	dataURL, err := r.Get("data_url")
	if err != nil {
		return Station{}, err
	}

	humanURL, err := r.Get("human_url")
	if err != nil {
		return Station{}, err
	}

	name, err := r.Get("name")
	if err != nil {
		return Station{}, err
	}

	riverName, err := r.Get("river_name")
	if err != nil {
		return Station{}, err
	}

	lat, err := r.Get("lat")
	if err != nil {
		return Station{}, err
	}

	lg, err := r.Get("lg")
	if err != nil {
		return Station{}, err
	}

	typeEnum, err := r.Get("type")
	if err != nil {
		return Station{}, err
	}

	unit, err := r.Get("unit")
	if err != nil {
		return Station{}, err
	}

	return Station{
		DataURL:   dataURL.(string),
		HumanURL:  humanURL.(string),
		Name:      name.(string),
		RiverName: riverName.(string),
		Lat:       lat.(float32),
		Lg:        lg.(float32),
		Type:      typeEnum.(goavro.Enum).Value,
		Unit:      unit.(string),
	}, nil
}

func recordToReading(r *goavro.Record) (Reading, error) {
	eventTime, err := r.Get("event_time")
	if err != nil {
		return Reading{}, err
	}

	value, err := r.Get("value")
	if err != nil {
		return Reading{}, err
	}

	return Reading{
		EventTime: time.Unix(eventTime.(int64), 0),
		Value:     value.(float32),
	}, nil
}

// Decode reads a Snapshot from AVRO binary format
func (s *Snapshot) Decode(reader io.Reader) error {
	decoded, err := snapshotCodec.Decode(reader)
	if err != nil {
		return err
	}

	r := decoded.(*goavro.Record)
	s.Station, err = recordToStation(r)
	if err != nil {
		return err
	}

	data, err := r.Get("readings")
	if err != nil {
		return err
	}
	innerRecords := data.([]interface{})
	for _, a := range innerRecords {
		u := a.(*goavro.Record)
		r, err := recordToReading(u)
		if err != nil {
			return err
		}

		s.Readings = append(s.Readings, r)
	}

	return nil
}

// Decode reads a Cache from AVRO binary format
func (c *Cache) Decode(reader io.Reader) error {
	decoded, err := cacheCodec.Decode(reader)
	if err != nil {
		return err
	}
	outerRecord := decoded.(*goavro.Record)
	data, err := outerRecord.Get("snapshots")
	if err != nil {
		return err
	}

	snapRecords := data.([]interface{})
	for _, a := range snapRecords {
		r := a.(*goavro.Record)
		s := Snapshot{
			ProcessingTime: time.Now(),
		}

		s.Station, err = recordToStation(r)
		if err != nil {
			return err
		}

		data, err := r.Get("readings")
		if err != nil {
			return err
		}
		innerRecords := data.([]interface{})
		for _, a := range innerRecords {
			u := a.(*goavro.Record)
			r, err := recordToReading(u)
			if err != nil {
				return err
			}

			s.Readings = append(s.Readings, r)
		}

		c.Add(&s)
	}

	return nil
}
