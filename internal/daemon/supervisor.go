package daemon

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rainchasers/report"
)

// Supervisor is a supervised instance of a running daemon
type Supervisor struct {
	*report.Logger

	doneC chan struct{}
	wg    sync.WaitGroup
}

// New creates a new daemon
func New(name string) *Supervisor {
	s := &Supervisor{
		Logger: createLogger(name),
		doneC:  make(chan struct{}),
	}

	s.Run(context.Background(), listenForTerminationSignal)
	return s
}

// Run executes a background function
//
// The daemon will shutdown if the function returns an unhandled error.
func (d *Supervisor) Run(ctx context.Context, fn func(ctx context.Context, d *Supervisor) error) {
	ctx = d.cancelOnClose(ctx)

	d.wg.Add(1)
	go func() {
		defer d.wg.Done()

		err := fn(ctx, d)
		if ctx.Err() == nil && err != nil {
			d.Action("daemon.interrupt", report.Data{
				"reason": "run process unhandled error",
				"error":  err.Error(),
			})
			d.Close()
		}
	}()
}

// CloseAfter closes the daemon after the specified delay
func (d *Supervisor) CloseAfter(duration time.Duration) {
	d.Run(context.Background(), func(ctx context.Context, d *Supervisor) error {
		timer := time.NewTimer(duration)
		defer timer.Stop()

		select {
		case <-timer.C:
			d.Close()
		case <-ctx.Done():
		}
		return nil
	})
}

// Wait blocks until the daemon has shutdown
func (d *Supervisor) Wait() {
	<-d.doneC

	timeout := 10 * time.Second
	c := make(chan struct{})
	go func() {
		defer close(c)
		d.wg.Wait()
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-c:
		<-d.Info("closed.ok", report.Data{})
	case <-timer.C:
		<-d.Action("closed.timeout", report.Data{})
	}

	// flush logs
	d.Logger.Close()
}

// Close invokes a deamon shutdown
func (d *Supervisor) Close() {
	// return early if already closing
	select {
	case <-d.doneC:
		return
	default:
	}

	// send close signal
	close(d.doneC)
}

// CloseWait invokes a deamon shutdown and blocks until complete
func (d *Supervisor) CloseWait() {
	d.Close()
	d.Wait()
}

func (d *Supervisor) cancelOnClose(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	d.wg.Add(1)
	go func() {
		<-d.doneC
		cancel()
		d.wg.Done()
	}()
	return ctx
}

func listenForTerminationSignal(ctx context.Context, d *Supervisor) error {
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	select {
	case s := <-sigC:
		d.Info("daemon.interrupt", report.Data{
			"reason": "user terminated",
			"signal": s,
		})
		d.Close()
	case <-ctx.Done():
	}

	return nil
}
