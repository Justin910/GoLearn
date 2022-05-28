package main

import (
	"context"
	"os"
	"time"
	"week03/transport"
)

type Option func(o *options)

type options struct {
	ctx  context.Context
	sigs []os.Signal

	stopTimeout time.Duration
	servers     []transport.Server
}

// WithServer with transport servers.
func WithServer(srv ...transport.Server) Option {
	return func(o *options) { o.servers = srv }
}

// WithSignal with exit signals.
func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// WithContext with service context.
func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// WithStopTimeout with app stop timeout.
func WithStopTimeout(t time.Duration) Option {
	return func(o *options) { o.stopTimeout = t }
}
