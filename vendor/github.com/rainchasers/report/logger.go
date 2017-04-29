package report

import (
	"encoding/json"
	"os"
	"time"
)

// Logger is the central logging agent on which to register events
type Logger struct {
	taskC  chan task
	stopC  chan bool
	global Data
	count  map[string]int
}

// Data is a string-keyed map of unstructured data relevant to the event
type Data map[string]interface{}

type command int

const (
	info command = iota
	action
	tock
	count
)

type task struct {
	command command
	event   string
	data    Data
	ackC    chan<- int
}

// New creates an instance of a logging agent
//
//     logger := report.New(report.Data{"service": "myAppName"})
//     defer logger.Stop()
//
func New(global Data) *Logger {
	logger := Logger{
		taskC:  make(chan task, 5),
		stopC:  make(chan bool),
		global: global,
		count:  make(map[string]int),
	}
	go logger.run()
	return &logger
}

// Info logs event that will provide context to any events requiring action.
//
//     logger := report.New(report.Data{"service": "myAppName"})
//     logger.Info("http.request", report.Data{"path":req.URL.Path, "ua":req.UserAgent()})
//
// If you would like to block based on the logline being processed, consume from the returned ack channel:
//
//     <-logger.Info("http.request", report.Data{"path":req.URL.Path, "ua":req.UserAgent()})
//
func (l *Logger) Info(event string, payload Data) <-chan int {
	ack := make(chan int)
	l.taskC <- task{
		command: info,
		event:   event,
		data:    payload,
		ackC:    ack,
	}
	return ack
}

// Action logs events that need intervention or resolving.
//
//     logger.Action("http.unavailable", report.Data{"path":req.URL.Path, "error":err.Error()})
//
// If you would like to block based on the logline being processed (for example before exiting),
// consume from the returned ack channel:
//
//     <-logger.Action("http.unavailable", report.Data{"path":req.URL.Path, "error":err.Error()})
//
func (l *Logger) Action(event string, payload Data) <-chan int {
	ack := make(chan int)
	l.taskC <- task{
		command: action,
		event:   event,
		data:    payload,
		ackC:    ack,
	}

	return ack
}

// Tick starts a timer event with a value for the later call to Tock
func (l *Logger) Tick() time.Time {
	return time.Now()
}

// Tock reports timer telemetry data recording the time since the Tick.
//
//     defer logger.Tock(report.Tick(), "mongo.query", report.Data{"q":query})
//
func (l *Logger) Tock(start time.Time, event string, payload Data) <-chan int {
	payload["ms"] = float64(time.Now().Sub(start).Nanoseconds()) / (float64(time.Millisecond) / float64(time.Nanosecond))

	ack := make(chan int)
	l.taskC <- task{
		command: tock,
		event:   event,
		data:    payload,
		ackC:    ack,
	}

	return ack
}

// Count returns the number of log events of a particular type since startup
//
//     logger := report.New(report.Data{"service": "myAppName"})
//     logger.Info("http.request", report.Data{})
//     logger.Info("http.request", report.Data{})
//     count := logger.Count("http.request")
//
func (l *Logger) Count(event string) int {
	ack := make(chan int)
	l.taskC <- task{
		command: count,
		event:   event,
		data:    Data{},
		ackC:    ack,
	}

	return <-ack
}

// RuntimeStatEvery log runtime stats at the specified interval
//
//     logger := report.New(report.Data{"service": "myAppName"})
//     logger.RuntimeStatEvery("runtime", time.Second*10)
//
func (l *Logger) RuntimeStatEvery(event string, duration time.Duration) {
	go func() {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

	statLoop:
		for {
			select {
			case <-ticker.C:
				l.Info(event, runtimeData())
			case <-l.stopC:
				break statLoop
			}
		}
	}()
}

// Stop shuts down the logging agent, further logging will result in a panic
//
//     logger := report.New(report.Data{"service": "myAppName"})
//     defer logger.Stop()
//
func (l *Logger) Stop() {
	close(l.taskC)
	close(l.stopC)
}

func (l *Logger) run() {
toNewTask:
	for t := range l.taskC {
		if t.command == count {
			n, exists := l.count[t.event]
			if exists {
				t.ackC <- n
			} else {
				t.ackC <- 0
			}
			close(t.ackC)
			continue toNewTask
		}

		n, exists := l.count[t.event]
		if exists {
			l.count[t.event] = n + 1
		} else {
			l.count[t.event] = 1
		}

		t.data["event"] = t.event
		t.data["timestamp"] = time.Now().Format(time.RFC3339Nano)
		for k, v := range l.global {
			t.data[k] = v
		}
		switch t.command {
		case info:
			t.data["type"] = "info"
		case action:
			t.data["type"] = "action"
		case tock:
			t.data["type"] = "timer"
		}

		os.Stdout.Write(map2Json(t.data))
		os.Stdout.Write([]byte("\n"))

		close(t.ackC)
	}
}

func map2Json(d Data) []byte {
	json, err := json.Marshal(d)
	if err != nil {
		return []byte(err.Error())
	}
	return json
}
