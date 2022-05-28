package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	opts   options
	ctx    context.Context
	cancel context.CancelFunc
	c      chan os.Signal
}

func NewApp(opts ...Option) *App {

	o := options{
		ctx:         context.Background(),
		sigs:        []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		stopTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)

	app := new(App)
	app.opts = o
	app.ctx = ctx
	app.cancel = cancel
	return app
}

func (app *App) Run() error {

	app.ctx, app.cancel = context.WithCancel(context.Background())

	egroup, ctx := errgroup.WithContext(app.ctx)

	wg := sync.WaitGroup{}
	for _, svr := range app.opts.servers {
		svr := svr
		egroup.Go(func() error {
			<-ctx.Done()
			log.Println("Server Recv Done Signal")
			stopCtx, cancel := context.WithTimeout(app.opts.ctx, app.opts.stopTimeout)
			defer cancel()
			return svr.Stop(stopCtx)
		})

		wg.Add(1)
		egroup.Go(func() error {
			log.Println("Distri Goroutine")
			wg.Done()
			return svr.Start(app.ctx)
		})
	}
	wg.Wait()
	log.Println("App Start Finish")

	app.c = make(chan os.Signal, 1)
	signal.Notify(app.c, app.opts.sigs...)
	egroup.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				log.Println("Recv Manual Stop Signal")
				return ctx.Err()
			case <-app.c:
				log.Println("Recv Kill Stop Signal")
				if err := app.Stop(); err != nil {
					return err
				}
			}
		}
	})
	if err := egroup.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
func (app *App) Stop() error {
	log.Println("Stop App")
	cancel := app.cancel
	if cancel != nil {
		cancel()
	}
	return nil
}

func (app *App) SimulateKillSignal() {
	app.c <- syscall.SIGINT
}
