package gauge

import (
	"io"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge/avro"
)

func typeToValue(t string) avro.TypeValues {
	switch t {
	case "level":
		return avro.Level
	case "flow":
		return avro.Flow
	case "temperature":
		return avro.Temperature
	case "rainfall":
		return avro.Rainfall
	}
	return avro.Level
}

func readingToAvro(r *Reading) *avro.Measure {
	m := avro.NewMeasure()
	m.Event_time = r.EventTime.Unix()
	m.Value = r.Value
	return m
}

func snapshotToAvro(s *Snapshot) *avro.Snapshot {
	a := avro.NewSnapshot()
	a.Data_url = s.Station.DataURL
	a.Alias_url = s.Station.AliasURL
	a.Human_url = s.Station.HumanURL
	a.Name = s.Station.Name
	a.River_name = s.Station.RiverName
	a.Lat = s.Station.Lat
	a.Lg = s.Station.Lg
	a.Type = typeToValue(s.Station.Type)
	a.Unit = s.Station.Unit
	a.Trace_id = s.TraceID
	a.Processing_time = s.ProcessingTime.Unix()

	for i := range s.Readings {
		a.Readings = append(a.Readings, readingToAvro(&s.Readings[i]))
	}

	return a
}

// Encode writes out the Snapshot in AVRO binary format
func (s *Snapshot) Encode(w io.Writer) error {
	a := snapshotToAvro(s)

	if err := a.Serialize(w); err != nil {
		return err
	}

	return nil
}

func avroToStation(a *avro.Snapshot) Station {
	return Station{
		DataURL:   a.Data_url,
		AliasURL:  a.Alias_url,
		HumanURL:  a.Human_url,
		Name:      a.Name,
		RiverName: a.River_name,
		Lat:       a.Lat,
		Lg:        a.Lg,
		Type:      a.Type.String(),
		Unit:      a.Unit,
	}
}

func avroToReading(a *avro.Measure) Reading {
	return Reading{
		EventTime: time.Unix(a.Event_time, 0),
		Value:     a.Value,
	}
}

// Decode reads a Snapshot from AVRO binary format
func (s *Snapshot) Decode(r io.Reader) error {
	a, err := avro.DeserializeSnapshot(r)
	if err != nil {
		return err
	}

	s.Station = avroToStation(a)
	for _, m := range a.Readings {
		s.Readings = append(s.Readings, avroToReading(m))
	}

	s.TraceID = a.Trace_id
	s.ProcessingTime = time.Unix(a.Processing_time, 0)

	return nil
}
