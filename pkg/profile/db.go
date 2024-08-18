package profile

import (
	"context"
	"time"
)

type SlowQueryLog struct {
	DBMS         string
	SlowQueryLog string
}

type options struct {
	Duration time.Duration
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithDuration(d time.Duration) Option {
	return optionFunc(func(o *options) {
		o.Duration = d
	})
}

type Profiler interface {
	Profiling(ctx context.Context, opts ...Option) ([]SlowQueryLog, error)
}
