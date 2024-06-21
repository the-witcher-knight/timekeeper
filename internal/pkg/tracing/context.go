package tracing

import (
	"context"
)

const (
	loggerCtxKey = "logger"
)

// SetInCtx add logger to context
func SetInCtx(ctx context.Context, l Tracer) context.Context {
	return context.WithValue(ctx, loggerCtxKey, l)
}

// FromCtx get logger from context
// if logger not exists in context returns a no-op Logger, that never writes out logs or internal errors
func FromCtx(ctx context.Context) Tracer {
	l, ok := ctx.Value(loggerCtxKey).(Tracer)
	if !ok {
		return NewNoop()
	}

	return l
}

func NewCtx(ctx context.Context) context.Context {
	newCtx := context.Background()
	return SetInCtx(newCtx, FromCtx(ctx))
}
