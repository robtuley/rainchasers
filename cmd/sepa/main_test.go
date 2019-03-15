package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"go.uber.org/goleak"
)

func TestDryRun(t *testing.T) {
	defer goleak.VerifyNoLeaks(t, goleak.IgnoreTopFunction("go.opencensus.io/stats/view.(*worker).start"))
	defer http.DefaultTransport.(*http.Transport).CloseIdleConnections()

	d := daemon.New("example")

	err := run(context.Background(), d)
	if err != nil {
		t.Fatal(err)
	}

	// validate log stream
	if d.Count("sepa.discover") != 1 {
		t.Fatal("sepa.discover event expected but not present")
	}
	if d.Count("sepa.updated") < 1 {
		t.Fatal("sepa.updated event expected but not present")
	}

	d.Close()
}
