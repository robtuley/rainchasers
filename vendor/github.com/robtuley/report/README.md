Logging Utility for Go
======================

An opinonated telemetry & logging utility for Go. 

[![GoDoc](https://godoc.org/github.com/robtuley/report?status.png)](https://godoc.org/github.com/robtuley/report)

+ one global logging stream per application
+ formatted as a stream of mostly unstructured JSON data events 
+ transport to an aggregator (e.g. Loggly)

Log Events
-----------

Log is a stream of data events that fall into 3 categories:

+ *Action*: an event that indicates a problem that needs attention to resolve. 
+ *Info*: an audit event adding context around any events requiring action. 
+ *Timer*: timing events that are to be used within visulation tools to better understand system dynamics.

Usage
-----

See `example/helloworld.go`:

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

Benchmark
---------

    go test -bench .