package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/rainchasers/report"
)

func TestRiverCacheCreationFromCatalogueURL(t *testing.T) {
	url := "https://app.rainchasers.com/catalogue.json"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log := report.New(ioutil.Discard, report.Data{})

	cache, err := NewRiverCache(ctx, url, log)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			t.Log("syntax error at byte offset", e.Offset)
		}
		t.Fatal(err)
	}
	if len(cache.Version) == 0 {
		t.Fatal("No version")
	}
	if len(cache.sectionMap) == 0 {
		t.Fatal("No sections!")
	}

	_, err = NewRiverCache(ctx, "https://app.rainchasers.com/a404.json", log)
	if err == nil {
		t.Fatal("No error despite 404 URL")
	}
}
