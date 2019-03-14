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
}
