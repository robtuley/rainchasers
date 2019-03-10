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

	doneC   chan struct{}
	doneFn  chan func()
	closedC chan struct{}
	wg      sync.WaitGroup
}

// New creates a new daemon
func New(name string) *Supervisor {
	doneFn := make(chan func())

	// cache done functions until channel closed
	// then execute them all
	go func() {
		fns := make([]func(), 0)
		for f := range doneFn {
			fns = append(fns, f)
		}
		for _, f := range fns {
			f()
		}
	}()

	s := &Supervisor{
		Logger:  createLogger(name),
		doneC:   make(chan struct{}),
		doneFn:  doneFn,
		closedC: make(chan struct{}),
	}

	go s.Run(context.Background(), listenForTerminationSignal)

	return s
}

// Run executes a job function
//
// The daemon will shutdown if the function returns an unhandled error.
func (d *Supervisor) Run(ctx context.Context, fn func(ctx context.Context, d *Supervisor) error) {
	ctx = d.cancelOnClose(ctx)

	d.wg.Add(1)
	err := fn(ctx, d)
	d.wg.Done()

	if ctx.Err() == nil && err != nil {
		d.Action("daemon.interrupt", report.Data{
			"reason": "run process unhandled error",
			"error":  err.Error(),
		})
		d.Close()
	}
}

// Done signals daemon shutdown
func (d *Supervisor) Done() <-chan struct{} {
	return d.doneC
}

// Close invokes a deamon shutdown
func (d *Supervisor) Close() {
	if d.isClosing() {
		// already closing, wait till complete
		<-d.closedC
		return
	}

	// send close signals
	close(d.doneC)
	close(d.doneFn)

	// wait for go-routines to exit cleanly (with timeout)
	c := make(chan struct{})
	go func() {
		defer close(c)
		d.wg.Wait()
	}()
	select {
	case <-c:
		<-d.Info("closed.ok", report.Data{})
	case <-time.After(10 * time.Second):
		<-d.Action("closed.timeout", report.Data{})
	}

	// flush logs
	d.Logger.Close()
	close(d.closedC)
}

// EndSpan writes the data from a completed trace span
func (d *Supervisor) EndSpan(ctx context.Context, err error, payload report.Data) <-chan int {
	if d.isClosing() {
		// do not end span if interrupted by closing
		ch := make(chan int)
		close(ch)
		return ch
	}

	return d.Logger.EndSpan(ctx, err, payload)
}

func (d *Supervisor) isClosing() bool {
	select {
	case <-d.doneC:
		return true
	default:
	}
	return false
}

func (d *Supervisor) cancelOnClose(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	d.doneFn <- cancel
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
