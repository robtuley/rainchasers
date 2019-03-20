package ea

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/report"
)

type stationListJson struct {
	Stations []stationJson `json:"items"`
}

type stationJson struct {
	Url              string `json:"@id"`
	RLOIid           string
	RLOIidRawJson    json.RawMessage `json:"RLOIid"`
	Name             string
	NameRawJson      json.RawMessage `json:"label"`
	RiverName        string
	RiverNameRawJson json.RawMessage `json:"riverName"`
	Lat              float32
	Lg               float32
	LatRawJson       json.RawMessage `json:"lat"`
	LgRawJson        json.RawMessage `json:"long"`
	Measures         []measureJson   `json:"measures"`
}

type measureJson struct {
	Url  string `json:"@id"`
	Name string `json:"qualifier"`
	Type string `json:"parameter"`
	Unit string `json:"unitName"`
}

// Discover finds all the available EA stations
func Discover(ctx context.Context) (map[string]gauge.Station, report.Span) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	url := "http://environment.data.gov.uk/flood-monitoring/id/stations"
	span := report.StartSpan("ea.discover").Field("url", url)
	stations := make(map[string]gauge.Station)

	resp, err := daemon.JSON(ctx, url)
	if err != nil {
		return stations, span.End(err)
	}
	defer resp.Body.Close()

	list := stationListJson{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&list)
	if err != nil {
		return stations, span.End(err)
	}

	for _, s := range list.Stations {
		// a known inconsistency is that the API can provide Lat, Lg or label as an array
		// so we use a defensive mechanism to parse these fields and let them be missing completely
		s.Lat, _ = daemon.ParseFloat(s.LatRawJson)
		s.Lg, _ = daemon.ParseFloat(s.LgRawJson)
		s.RLOIid, _ = daemon.ParseString(s.RLOIidRawJson)
		s.Name, _ = daemon.ParseString(s.NameRawJson)
		s.RiverName, _ = daemon.ParseString(s.RiverNameRawJson)

		for _, m := range s.Measures {

			switch m.Type {
			case "flow", "level", "temperature", "rainfall":
			default:
				continue
			}

			// if a RLOIid (river levels on internet ID) is
			// available, use this as a AliasID otherwise use
			// the measure URL
			aliasURL := m.Url
			if s.RLOIid != "" {
				aliasURL = "rloi://" + s.RLOIid
			}

			// no snapshot readings are available
			// s.DateTime and s.Value left as defaults

			station := gauge.Station{
				DataURL:   m.Url,
				AliasURL:  aliasURL,
				HumanURL:  s.Url,
				Name:      s.Name,
				RiverName: s.RiverName,
				Lat:       s.Lat,
				Lg:        s.Lg,
				Type:      m.Type,
				Unit:      m.Unit,
			}
			stations[m.Url] = station
		}
	}

	span = span.Field("stations_count", len(stations))
	return stations, span.End()
}
