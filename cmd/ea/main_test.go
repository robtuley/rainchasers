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
	// if d.Count("snapshot.published") < 1 {
	//	t.Fatal("No snapshot.published events")
	// }
	d.Close()
}
