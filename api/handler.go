package main

import (
	"net/http"

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
	defer r.Body.Close()
	if h.IsReady {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

// Export exports cache state as an avro document
func (h *Handler) Export(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "avro/binary")
	if err := h.Gauge.Encode(w); err != nil {
		<-h.Log.Action("response.cache.error", report.Data{"error": err.Error()})
		return
	}
}

// API handles public data API requests
func (h *Handler) API(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch {
	case r.URL.Path == "/":
		h.Docs(w, r)
	case r.URL.Path == "/events":
		w.Write([]byte("event stream"))
	default:
		w.Write([]byte("individual item UUID"))
	}
	//w.Header().Set("Content-Type", "application/json")
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
