package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rainchasers/content/internal/gauge"
	"github.com/rainchasers/report"
)

type customTime struct {
	time.Time
}

const ctLayout = time.RFC3339

var nilTime = (time.Time{}).UnixNano()

func (ct *customTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *customTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

func (ct *customTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type recentJSON []struct {
	ID     int    `json:"location"`
	Title  string `json:"titleEN"`
	Coords struct {
		Lat float32 `json:"latitude"`
		Lng float32 `json:"longitude"`
	} `json:"coordinates"`
	Parameters []struct {
		Name        string     `json:"paramNameEN"`
		LatestValue float32    `json:"latestValue"`
		LatestTime  customTime `json:"latestTime"`
		Units       string     `json:"units"`
	} `json:"parameters"`
}

const recentURL = "https://api.naturalresources.wales/rivers-and-seas/v1/api/StationData"
const recentKeyHeader = "Ocp-Apim-Subscription-Key"

// Recent fetches recent NRW readings
func recent(ctx context.Context, apiKey string) ([]gauge.Snapshot, report.Span) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	span := report.StartSpan("nrw.recent").Field("url", recentURL)

	// do the request
	req, err := http.NewRequest("GET", recentURL, nil)
	if err != nil {
		return nil, span.End(err)
	}
	req.WithContext(ctx)
	req.Header.Add("Accept", "application/json")
	req.Header.Set(recentKeyHeader, apiKey)

	client := &http.Client{
		Timeout: time.Second * 60,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, span.End(err)
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		msg := "Status code " + strconv.Itoa(resp.StatusCode) + " : "
		bb, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			msg = msg + string(bb)
		}
		return nil, span.End(errors.New(msg))
	}

	// parse response
	snaps, err := parseRecent(resp.Body)
	span = span.Field("snapshots_count", len(snaps))
	return snaps, span.End(err)
}

func parseRecent(r io.Reader) ([]gauge.Snapshot, error) {
	snaps := make([]gauge.Snapshot, 0)

	parsed := recentJSON{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&parsed)
	if err != nil {
		return snaps, err
	}

	for _, feature := range parsed {
		if len(feature.Parameters) != 1 {
			continue
		}
		parameters := feature.Parameters[0]
		rtoi := strconv.Itoa(feature.ID)

		var stationType string
		switch parameters.Name {
		case "Rainfall":
			stationType = "rainfall"
		case "River Level":
			stationType = "level"
		default:
			continue
		}

		station := gauge.Station{
			DataURL:   "rloi://" + rtoi,
			AliasURL:  "rloi://" + rtoi,
			HumanURL:  "https://rloi.naturalresources.wales/ViewDetails?station=" + rtoi,
			Name:      feature.Title,
			RiverName: "", // not available
			Lat:       feature.Coords.Lat,
			Lg:        feature.Coords.Lng,
			Type:      stationType,
			Unit:      parameters.Units,
		}

		reading := gauge.Reading{
			EventTime: parameters.LatestTime.Time,
			Value:     parameters.LatestValue,
		}

		snaps = append(snaps, gauge.Snapshot{
			Station:  station,
			Readings: []gauge.Reading{reading},
		})
	}

	return snaps, nil
}
