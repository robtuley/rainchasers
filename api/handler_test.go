package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/report"
)

func TestUUIDHandlerRespondsWithSectionAndMeasures(t *testing.T) {
	url := "https://app.rainchasers.com/catalogue.json"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log := report.New(ioutil.Discard, report.Data{})
	riverCache, err := NewRiverCache(ctx, url, log)
	if err != nil {
		t.Fatal(err)
	}
	gaugeCache := gauge.NewCache(ctx, time.Hour)

	h := Handler{
		Log:     report.New(ioutil.Discard, report.Data{}),
		Rivers:  riverCache,
		Gauge:   gaugeCache,
		IsReady: true,
	}

	req, err := http.NewRequest("GET", "/338189e7-5af2-4b4b-964d-324f22c7f259", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	http.HandlerFunc(h.UUID).ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check the response has a body
	if len(rr.Body.String()) == 0 {
		t.Fatal("No body returned", rr.Body.String())
	}
}
