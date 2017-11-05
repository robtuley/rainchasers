package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/report"
)

// Handler has methods to handle internal and external API requests
type Handler struct {
	Log     *report.Logger
	Gauge   *gauge.Cache
	Rivers  *RiverCache
	IsReady bool
}

// Readiness handles internal k8s readiness probes
func (h *Handler) Readiness(w http.ResponseWriter, r *http.Request) {
	if h.IsReady {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

// Export exports cache state as an avro document
func (h *Handler) Export(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "avro/binary")
	if err := h.Gauge.Encode(w); err != nil {
		<-h.Log.Action("response.cache.error", report.Data{"error": err.Error()})
		return
	}
}

// API handles public data API requests
func (h *Handler) API(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		h.Explore(w, r)
	case r.URL.Path == "/events":
		w.Write([]byte("event stream"))
	default:
		h.UUID(w, r)
	}
}

// Explore handles public data API documentation
func (h *Handler) Explore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	writeCatalogueHTML(h, w)
}

// UUID responds with a single entity
func (h *Handler) UUID(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Trim(r.URL.Path, "/")

	s, ok := h.Rivers.Load(uuid)
	if !ok {
		http.NotFound(w, r)
		return
	}

	for k := range s.Measures {
		g, ok := h.Gauge.Load(s.Measures[k].DataURL)
		if ok {
			s.Measures[k].Snapshot = &g
		} else {
			h.Log.Action("gauge.cache.miss", report.Data{
				"river_uuid": uuid,
				"data_url":   s.Measures[k].DataURL,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
