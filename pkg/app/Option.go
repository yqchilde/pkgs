package app

import (
	"context"
	"os"

	"github.com/yqchilde/gint/pkg/logger"
	"github.com/yqchilde/gint/pkg/transport"
)

type Option func(o *options)

type options struct {
	name    string
	sigs    []os.Signal
	ctx     context.Context
	log     logger.Logger
	servers []transport.Server
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

func WithLogger(logger logger.Logger) Option {
	return func(o *options) {
		o.log = logger
	}
}

func Server(srv ...transport.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}
