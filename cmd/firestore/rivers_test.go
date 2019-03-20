package main

import (
	"context"
	"encoding/json"
	"testing"
)

func TestRiverCacheCreationFromCatalogueURL(t *testing.T) {
	url := "https://app.rainchasers.com/catalogue.json"

	cache := NewRiverCache(url)
	_, span := cache.Update(context.Background())
	if err := span.Err(); err != nil {
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

	cache = NewRiverCache("https://app.rainchasers.com/a404.json")
	_, span = cache.Update(context.Background())
	if span.Err() == nil {
		t.Fatal("No error despite 404 URL")
	}
}
