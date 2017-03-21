Logging Utility for Go
======================

An opinonated telemetry & logging utility for Go. 

[![GoDoc](https://godoc.org/github.com/robtuley/report?status.png)](https://godoc.org/github.com/robtuley/report)

[![Code Climate](https://codeclimate.com/github/robtuley/report/badges/gpa.svg)](https://codeclimate.com/github/robtuley/report)

[![Issue Count](https://codeclimate.com/github/robtuley/report/badges/issue_count.svg)](https://codeclimate.com/github/robtuley/report)

+ one global logging stream per application
+ format as unstructured JSON data events 
+ transport to an aggregator (e.g. Loggly)

Log Events
-----------

Log is a stream of data events that fall into 3 categories:

+ *Action*: an event that indicates a problem that needs attention to resolve. 
+ *Info*: an audit event adding context around any events requiring action. 
+ *Timer*: timing events that are to be used within visulation tools to better understand system dynamics.

Usage
-----

See `example/webserver/helloworld.go`:

```
func main() {
	defer report.Drain()		
	report.StdOut()
	// OR more likely:
	// report.BatchPostToUrl("yourLogglyBulkUploadUrl")

	// add data for all log events if mixed aggregation
	report.Global(report.Data{"service": "myAppName"})
	
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

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// an *actionable* event that needs resolution
		report.Action("http.listening.fail", report.Data{"error": err.Error()})
	}
}
```

Observing the Log Stream
------------------------

If your application needs to hook into the log stream (for testing the output log API, or alternative processing) a `Observer` type function can be passed to the `report.Observe` method. This will be invoked with each log line. Make sure you contain any logging within this function to prevent circular recursion!

See `example/daemon/testable.go` for a full example:

```
// use observer to accumulate log messages
logC := make(chan report.Data, 50)
log2Channel := func(d report.Data) {
	select {
	case logC <- d:
	default:
	// (a non-blocking publish)
	}
}
report.Observe(log2Channel)
``` 

Back-pressure on Pipeline Overflow
----------------------------------

    go test -bench .