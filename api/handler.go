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
		h.UUID(strings.Trim(r.URL.Path, "/"), w, r)
	}
}

// Explore handles public data API documentation
func (h *Handler) Explore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	writeCatalogueHTML(h, w)
}

// UUID responds with a single entity
func (h *Handler) UUID(uuid string, w http.ResponseWriter, r *http.Request) {
	s, ok := h.Rivers.Load(uuid)
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(s)
}
