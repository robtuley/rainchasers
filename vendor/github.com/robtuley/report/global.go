package report

import (
	"time"
)

// Global adds a global data element to all logged events (e.g. if aggregating across services).
//
//     report.Global(report.Data{"service":"MyServiceName","version":"1.5.2"})
//
func Global(payload Data) {
	channel.AddGlobal <- payload
}

// Observe allows a custom callback into the log stream.
//
//     report.Observe(func (d report.Data){ log.Println(d) })
//
func Observe(o Observer) {
	channel.AddObserver <- o
}

// MaxObserverConcurrency specifies the maximum concurrent executing observer go-routines
const MaxObserverConcurrency = 50

func init() {
	go func() {
		semC := make(chan bool, MaxObserverConcurrency)
		observerExe := func(f Observer, e Data) {
			f(e)
			<-semC
		}

		globals := Data{}
		var observers []Observer

		for {
			select {
			case evt, more := <-channel.RawEvents:
				if !more {
					// ensure all observers have completed execution
					// by filling semaphore channel before allowing
					// drain completion
					for i := 0; i < cap(semC); i++ {
						semC <- true
					}
					channel.DrainSignal <- true
					return
				}
				evt["timestamp"] = time.Now().Format(time.RFC3339Nano)
				for k, v := range globals {
					evt[k] = v
				}
				channel.WithGlobals <- evt
				for _, f := range observers {
					// max concurrency is governed by trying to publish
					// to the buffered semaphore channel before spawning
					// another go-routine
					semC <- true
					go observerExe(f, evt)
				}
			case g := <-channel.AddGlobal:
				for k, v := range g {
					globals[k] = v
				}
			case o := <-channel.AddObserver:
				observers = append(observers, o)
			}
		}
	}()
}
