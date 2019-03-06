package daemon

import (
	"os"
	"time"

	"github.com/rainchasers/report"
)

// Responds to following env vars:
//   HONEYCOMB_API_KEY (no default, blank to skip honeycomb integration)
//   NODE_NAME (no default, name of k8s node for logging)
//   POD_NAME (no default, name of k8s pod for logging)
func createLogger(name string) *report.Logger {
	honeycombKey := os.Getenv("HONEYCOMB_API_KEY")
	nodeName := os.Getenv("NODE_NAME")
	podName := os.Getenv("POD_NAME")

	log := report.New(name)
	log.Export(report.StdOutJSON())
	if len(honeycombKey) > 0 {
		log.Export(report.Honeycomb(honeycombKey, "rainchasers"))
	}
	log.RuntimeEvery(time.Minute)
	if len(nodeName) > 0 {
		log.Baggage("node", nodeName)
	}
	if len(podName) > 0 {
		log.Baggage("pod", podName)
	}

	return log
}
