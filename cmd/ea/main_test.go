package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"go.uber.org/goleak"
)

func TestDryRun(t *testing.T) {
	defer goleak.VerifyNoLeaks(t, goleak.IgnoreTopFunction("go.opencensus.io/stats/view.(*worker).start"))
	defer http.DefaultTransport.(*http.Transport).CloseIdleConnections()

	// define an accelerated dry run
	cfg := config{
		RefreshPeriodInSeconds: 3,
	}

	// init daemon supervisor & context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	d := daemon.New("example")

	// perform the dry run
	err := cfg.run(ctx, d)
	if err != nil {
		t.Fatal(err)
	}

	// validate the log interface
	if d.Count("snapshot.published") < 1 {
		t.Fatal("No snapshot.published events")
	}
	if err := d.Err(); err != nil {
		t.Fatal(err)
	}
	d.Close()
}
