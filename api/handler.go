package main

import (
	"context"
	"net/http"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/report"
)

// Handler has methods to handle internal and external API requests
type Handler struct {
	Log           *report.Logger
	Cache         *gauge.Cache
	IsReady       bool
	ClientTimeout time.Duration
}

// Readiness handles internal k8s readiness probes
func (h Handler) Readiness(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if h.IsReady {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

// Export exports cache state as an avro document
func (h Handler) Export(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "avro/binary")
	if err := h.Cache.Encode(w); err != nil {
		<-h.Log.Action("response.cache.error", report.Data{"error": err.Error()})
		return
	}
}

// API handles public data API requests
func (h Handler) API(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	// internal timeout 90% of client timeout to allow clean return
	timeout := time.Duration(h.ClientTimeout.Nanoseconds()*9/10) * time.Nanosecond
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	// TODO: do actual API
	_ = ctx
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{status:ok}"))
}
