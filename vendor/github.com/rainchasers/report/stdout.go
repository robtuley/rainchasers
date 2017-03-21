package report

import (
	"os"
)

// StdOut write events to stdout
func StdOut() {
	go stdoutWriter()
}

func stdoutWriter() {
	for {
		json, more := <-channel.JSONEncoded
		if !more {
			channel.DrainSignal <- true
			return
		}
		os.Stdout.Write([]byte(json + "\n"))
	}
}
