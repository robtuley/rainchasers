package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rainchasers/content/internal/daemon"
	"go.uber.org/goleak"
)

func TestDryRun(t *testing.T) {
	// verify no goroutine leaks
	defer goleak.VerifyNoLeaks(t, goleak.IgnoreTopFunction("go.opencensus.io/stats/view.(*worker).start"))
	defer http.DefaultTransport.(*http.Transport).CloseIdleConnections()

	// define an accelerated dry run
	cfg := config{
		RefreshPeriodInSeconds: 5,
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
	expected := []string{
		"topic.connected",
		"snapshot.published",
	}
	for _, evt := range expected {
		if d.Count(evt) == 0 {
			t.Fatal(evt + " event expected but not present")
		}
	}

	if err := d.Err(); err != nil {
		t.Fatal(err)
	}
	d.CloseWait()
}
