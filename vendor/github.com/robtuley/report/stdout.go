package report

import (
	"os"
)

// Write all recorded events to stdout
func StdOut() {
	go stdoutWriter()
}

func stdoutWriter() {
	for {
		json, more := <-channel.JsonEncoded
		if !more {
			channel.DrainSignal <- true
			return
		}
		os.Stdout.Write([]byte(json + "\n"))
	}
}
