package main

import (
	"bytes"
	"testing"
)

func TestRecentParse(t *testing.T) {
	b := bytes.NewBufferString(jsonResponseFromAPI)
	snaps, err := parseRecent(b)
	if err != nil {
		t.Fatal(err)
	}

	if len(snaps) == 0 {
		t.Fatal("No snaps parsed")
	}

	for _, snap := range snaps {
		s := snap.Station

		if s.DataURL == "rloi://0" {
			t.Fatal("zero RTOI detected", s)
		}
		if len(s.Name) < 3 {
			t.Fatal("Missing name detected ", s)
		}
		if len(s.Unit) < 1 {
			t.Fatal("Missing unit detected ", s)
		}

		if len(snap.Readings) != 1 {
			t.Fatal("Missing reading ", snap)
		}
	}
}
