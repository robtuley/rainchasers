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

	ctx       context.Context
	cancelCtx func()
	wg        sync.WaitGroup
	closeOnce sync.Once
}

// New creates a new daemon
func New(name string, timeout time.Duration) *Supervisor {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	s := &Supervisor{
		Logger:    createLogger(name),
		ctx:       ctx,
		cancelCtx: cancel,
	}

	s.Supervise(listenForTerminationSignal)

	return s
}

// Context provides the timeout context of the daemon
func (d *Supervisor) Context() context.Context {
	return d.ctx
}

// Supervise executes a background func (non-blocking)
//
// If the function passed returns an error the full daemon will
// be shutdown as the error is regarded as unhandled.
func (d *Supervisor) Supervise(fn func(d *Supervisor) error) {
	d.wg.Add(1)
	go func() {
		err := fn(d)
		d.wg.Done()
		if err != nil {
			d.Action("daemon.interrupt", report.Data{
				"reason": "supervised process unhandled error",
				"error":  err.Error(),
			})
			d.Close()
		}
	}()
}

// Run executes the primary daemon function
//
// The daemon will shutdown if the function returns an
// unhandled error.
func (d *Supervisor) Run(fn func(d *Supervisor) error) {
	d.wg.Add(1)
	err := fn(d)
	d.wg.Done()

	if err != nil {
		d.Action("daemon.interrupt", report.Data{
			"reason": "run process unhandled error",
			"error":  err.Error(),
		})
		d.Close()
	}
}

// Close invokes a full blocking deamon shutdown
//
// Be careful not to invoke a WG deadlock. If you are inside the
// a main runner, don't block execution completion of the runner
// as shutdown will wait for it to complete
func (d *Supervisor) Close() {
	d.closeOnce.Do(func() {
		// invoke daemon context cancellation
		d.cancelCtx()

		// wait for 10 seconds for routines to close down cleanly
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		c := make(chan bool)
		go func() {
			defer close(c)
			d.wg.Wait()
		}()
		select {
		case <-c:
			<-d.Info("daemon.stopped.ok", report.Data{})
		case <-ctx.Done():
			<-d.Action("daemon.stopped.timeout", report.Data{})
		}
		cancel()

		// flush logs and close logger
		d.Logger.Close()
	})
}

// EndSpan writes the data from a completed trace span
func (d *Supervisor) EndSpan(ctx context.Context, err error, payload report.Data) <-chan int {
	if d.ctx.Err() == nil {
		// end span only if not interrupted by shutdown
		return d.Logger.EndSpan(ctx, err, payload)
	}

	ch := make(chan int)
	close(ch)
	return ch
}

func listenForTerminationSignal(d *Supervisor) error {
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
	case <-d.ctx.Done():
	}

	return nil
}
