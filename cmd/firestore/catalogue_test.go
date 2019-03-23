package main

import (
	"context"
	"encoding/json"
	"testing"
)

func TestParseCatalogue(t *testing.T) {
	sections, span := parseCatalogue(context.Background())
	if err := span.Err(); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			t.Log("syntax error at byte offset", e.Offset)
		}
		t.Fatal(err)
	}
	if len(sections) == 0 {
		t.Fatal("No sections!")
	}
}
