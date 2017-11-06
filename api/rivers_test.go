package main

import (
	"context"
	"encoding/json"
	"testing"
)

func TestRiverCacheCreationFromCatalogueURL(t *testing.T) {
	url := "https://app.rainchasers.com/catalogue.json"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache, err := NewRiverCache(ctx, url)
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

	_, err = NewRiverCache(ctx, "https://app.rainchasers.com/a404.json")
	if err == nil {
		t.Fatal("No error despite 404 URL")
	}
}
