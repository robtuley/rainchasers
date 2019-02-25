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
	Log     *report.Logger
	Context context.Context

	wg            sync.WaitGroup
	cancelContext func()
	shutdownOnce  sync.Once
	err           error
	errMutex      sync.Mutex
}

// New creates a new daemon
func New(name string, timeout time.Duration) *Supervisor {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	s := &Supervisor{
		Log:           createLogger(name),
		Context:       ctx,
		cancelContext: cancel,
	}

	s.Supervise(listenForTerminationSignal)

	return s
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
			d.storeError(err)
			d.Log.Action("daemon.interrupt", report.Data{
				"reason": "supervised process unhandled error",
				"error":  err.Error(),
			})
			d.Shutdown()
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
		d.storeError(err)
		d.Log.Action("daemon.interrupt", report.Data{
			"reason": "run process unhandled error",
			"error":  err.Error(),
		})
		d.Shutdown()
	}
}

// Shutdown invokes a full blocking deamon shutdown
//
// Be careful not to invoke a WG deadlock. If you are inside the
// a main runner, don't block execution completion of the runner
// as shutdown will wait for it to complete
func (d *Supervisor) Shutdown() {
	d.shutdownOnce.Do(func() {
		// invoke daemon context cancellation
		d.cancelContext()

		// wait for 10 seconds for routines to close down cleanly
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		c := make(chan bool)
		go func() {
			defer close(c)
			d.wg.Wait()
		}()
		select {
		case <-c:
			<-d.Log.Info("daemon.stopped.ok", report.Data{})
		case <-ctx.Done():
			<-d.Log.Action("daemon.stopped.timeout", report.Data{})
		}
		cancel()

		// flush logs and close cleanly
		if err := d.Log.Err(); err != nil {
			d.storeError(err)
		}
		d.Log.Close()
	})
}

// Err returns the first error on an unclean shutdown
func (d *Supervisor) Err() error {
	d.errMutex.Lock()
	defer d.errMutex.Unlock()

	return d.err
}

func (d *Supervisor) storeError(err error) {
	d.errMutex.Lock()
	defer d.errMutex.Unlock()

	if d.err == nil {
		d.err = err
	}
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
		d.Log.Info("daemon.interrupt", report.Data{
			"reason": "user terminated",
			"signal": s,
		})
		d.Shutdown()
	case <-d.Context.Done():
	}

	return nil
}
