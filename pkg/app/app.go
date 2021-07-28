package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/yqchilde/gint/pkg/config"
	"github.com/yqchilde/gint/pkg/logger"
)

type App struct {
	cfg    *config.Config
	opts   *options
	ctx    context.Context
	cancel func()
	log    logger.Logger
}

func New(cfg *config.Config, opts ...Option) *App {
	options := &options{
		ctx:  context.Background(),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		log:  logger.GetLogger(),
	}
	for _, o := range opts {
		o(options)
	}

	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		cfg:    cfg,
		opts:   options,
		ctx:    ctx,
		cancel: cancel,
		log:    logger.GetLogger(),
	}
}

func (a *App) Run() error {
	a.log.Infof("app_name: %s", a.opts.name)
	eg, ctx := errgroup.WithContext(a.ctx)

	// start server
	for _, srv := range a.opts.servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		eg.Go(func() error {
			return srv.Start()
		})
	}

	// watch signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, a.opts.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s := <-quit:
				a.log.Infof("receive a quit signal: %s", s.String())
				return a.Stop()
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
