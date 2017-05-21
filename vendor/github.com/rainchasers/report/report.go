package report

import (
	"encoding/json"
	"errors"
	"io"
	"sync"
	"time"
)

// Logger is the central logging agent on which to register events
type Logger struct {
	writer         io.Writer
	taskC          chan task
	stopC          chan struct{}
	global         Data
	count          map[string]int
	lastError      error
	lastErrorMutex sync.Mutex
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

const msPerNs = float64(time.Millisecond) / float64(time.Nanosecond)

// New creates an instance of a logging agent
//
//     log := report.New(os.Stdout, report.Data{"service": "myAppName"})
//     defer logger.Stop()
//
func New(writer io.Writer, global Data) *Logger {
	logger := Logger{
		writer: writer,
		taskC:  make(chan task, 1),
		stopC:  make(chan struct{}),
		global: global,
		count:  make(map[string]int),
	}
	go logger.run()
	return &logger
}

// Info logs event that provide telemetry measures or context to any events requiring action.
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

// Action events that need intervention or resolving.
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
	payload["ms"] = float64(time.Now().Sub(start).Nanoseconds()) / msPerNs

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
//     log := report.New(os.Stdout, report.Data{"service": "myAppName"})
//     log.RuntimeStatEvery("runtime", time.Second*10)
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

// LastError returns the last Actionable log event or encoding error if either occurred
func (l *Logger) LastError() error {
	l.lastErrorMutex.Lock()
	defer l.lastErrorMutex.Unlock()

	return l.lastError
}

// Stop shuts down the logging agent, further logging will result in a panic
//
//     log := report.New(os.Stdout, report.Data{"service": "myAppName"})
//     defer log.Stop()
//
func (l *Logger) Stop() {
	close(l.taskC)
	close(l.stopC)
}

func (l *Logger) run() {
	encoder := json.NewEncoder(l.writer)

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
			l.lastErrorMutex.Lock()
			l.lastError = errors.New("Actionable event: " + t.event)
			l.lastErrorMutex.Unlock()
		case tock:
			t.data["type"] = "timer"
		}

		if err := encoder.Encode(t.data); err != nil {
			l.lastErrorMutex.Lock()
			l.lastError = err
			l.lastErrorMutex.Unlock()
		}
		close(t.ackC)
	}
}
