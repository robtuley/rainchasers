// Example webserver demonstrating report package logging
package main

import (
	"github.com/robtuley/report"
	"io"
	"net/http"
	"time"
)

func main() {
	defer report.Drain()
	report.StdOut()
	// OR more likely:
	// report.BatchPostToUrl("yourLogglyBulkUploadUrl")

	// add data for all log events if mixed aggregation
	report.Global(report.Data{"service": "myAppName"})

	// report runtime stats evry 10s
	report.RuntimeStatsEvery(time.Second * 10)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// timer to record response time and details
		defer report.Tock(report.Tick(), "http.response", report.Data{
			"host": req.URL.Host,
			"path": req.URL.Path,
			"ua":   req.UserAgent(),
		})
		io.WriteString(res, "Hello World")
	})

	// an *info* event to provide context
	report.Info("http.listening", report.Data{"port": 8080})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			// an *actionable* event that needs resolution
			report.Action("http.listening.fail", report.Data{"error": err.Error()})
		}
	}()

	// to demo drain exist after 30 seconds
	<-time.After(time.Second * 60)
}
