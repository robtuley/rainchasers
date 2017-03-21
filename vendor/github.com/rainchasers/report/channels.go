package report

import (
	"log"
)

// Data is a string-keyed map of unstructured data relevant to the event
type Data map[string]interface{}

// Observer allows custom functional interception of the log stream
type Observer func(Data)

//     info.go, action.go, timer.go
// --> channel.RawEvents
// --> global.go <-- channel.AddGlobal
//               --> channel.Drain
// --> channel.WithGlobals
// --> json.go --> channel.Drain
// --> channel.JSONEncoded
// --> broadcast.go --> channel.Drain
var channel struct {
	RawEvents   chan Data
	WithGlobals chan Data
	AddGlobal   chan Data
	AddObserver chan Observer
	JSONEncoded chan string
	DrainSignal chan bool
	IsDraining  chan bool
}

func init() {
	channel.RawEvents = make(chan Data, 50)
	channel.WithGlobals = make(chan Data, 50)
	channel.AddGlobal = make(chan Data)
	channel.AddObserver = make(chan Observer)
	channel.JSONEncoded = make(chan string, 50)
	channel.DrainSignal = make(chan bool)
	channel.IsDraining = make(chan bool)
}

// Drain waits for events to drain down before exiting, usually called before exit on main func
//
//	func main {
//		defer report.Drain()
//
//		// ... snip ...
//	}
func Drain() {
	select {
	case <-channel.IsDraining:
		return
	default:
	}
	close(channel.IsDraining)

	close(channel.RawEvents)
	<-channel.DrainSignal

	close(channel.WithGlobals)
	<-channel.DrainSignal

	close(channel.JSONEncoded)
	<-channel.DrainSignal
}

func publishRawEvent(payload Data) {
	select {
	case <-channel.IsDraining:
		log.Println("discarded:>", payload)
	default:
		channel.RawEvents <- payload
	}
}
