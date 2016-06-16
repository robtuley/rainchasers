package main

import (
	"github.com/robtuley/report"
)

func bufferLogStream(bufferSize int) <-chan report.Data {
	logC := make(chan report.Data, bufferSize)
	log2Channel := func(d report.Data) {
		select {
		case logC <- d:
		default:
			// (a non-blocking publish)
		}
	}
	report.Observe(log2Channel)
	return logC
}

func validateLogStream(logC <-chan report.Data) error {
	var n int

Consume:
	for {
		select {
		case <-logC:
			n = n + 1
		default:
			break Consume
		}
	}

	report.Info("daemon.validation", report.Data{"buffer_size": n})
	return nil
}
