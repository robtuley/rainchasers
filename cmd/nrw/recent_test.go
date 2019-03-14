package main

import (
	"bytes"
	"testing"
)

func TestRecentParse(t *testing.T) {
	b := bytes.NewBufferString(jsonResponseFromApi)
	snaps, err := parseRecent(b)
	if err != nil {
		t.Fatal(err)
	}

	if len(snaps) == 0 {
		t.Fatal("No snaps parsed")
	}
}
