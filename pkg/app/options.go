package app

import (
	"context"
	"os"
	"time"

	"github.com/yqchilde/gin-skeleton/pkg/log"
	"github.com/yqchilde/gin-skeleton/pkg/transport"
)

type Option func(o *options)

type options struct {
	id   string
	name string

	sigs []os.Signal
	ctx  context.Context

	logger           log.Logger
	registrarTimeout time.Duration
	servers          []transport.Server
}

func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func Server(srv ...transport.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}
