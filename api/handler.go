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
		h.Docs(w, r)
	case r.URL.Path == "/catalogue":
		h.Catalogue(w, r)
	case r.URL.Path == "/events":
		w.Write([]byte("event stream"))
	default:
		h.UUID(strings.Trim(r.URL.Path, "/"), w, r)
	}
}

// Docs handles public data API documentation
func (h *Handler) Docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
<html>
<head>
 <title>Rainchasers Data</title>
</head>
<body>
 <p>This is the documentation/explore area		
	`))
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

// Catalogue responds with full point catalogue
func (h *Handler) Catalogue(w http.ResponseWriter, r *http.Request) {
	type point struct {
		UUID string  `json:"uuid"`
		Name string  `json:"name"`
		URL  string  `json:"url"`
		Lat  float32 `json:"lat"`
		Lng  float32 `json:"lng"`
	}

	pts := make([]point, 0, h.Rivers.Count())
	h.Rivers.Each(func(s Section) bool {
		pts = append(pts, point{
			UUID: s.UUID,
			Name: s.SectionName + ", " + s.RiverName,
			URL:  "todo",
			Lat:  s.Putin.Lat,
			Lng:  s.Putin.Lng,
		})
		return true
	})

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(pts)
}
