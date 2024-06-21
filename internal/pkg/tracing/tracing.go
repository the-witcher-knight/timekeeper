package tracing

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/logging"
)

type Tracer interface {
	Info(ctx context.Context, msg string, args ...Field)
	Debug(ctx context.Context, msg string, args ...Field)
	Error(ctx context.Context, err error, msg string, withStack bool, args ...Field)
	With(args ...Field) Tracer
}

type tracer struct {
	logger      logging.Logger
	extraFields []Field
}

func NewTracer(logger logging.Logger, extraFields ...Field) Tracer {
	return tracer{
		logger:      logger,
		extraFields: extraFields,
	}
}

func (t tracer) Info(ctx context.Context, msg string, args ...Field) {
	t.logger.Info(msg, toZapFieldSlice(args)...)
}

func (t tracer) Debug(ctx context.Context, msg string, args ...Field) {
	t.logger.Debug(msg, toZapFieldSlice(args)...)
}

func (t tracer) Error(ctx context.Context, err error, msg string, withStack bool, args ...Field) {
	args = append(args, String("exception.full_message", err.Error()))

	originErr := errors.GetRootErr(err)
	if originErr != nil {
		args = append(args,
			String("exception.message", originErr.Error()),
			String("exception.type", fmt.Sprintf("%T", originErr)),
		)
	}

	if withStack {
		stackTraced, ok := err.(errors.StackTracer)
		if ok {
			args = append(args, String("exception.stack", fmt.Sprintf("%v", stackTraced.StackTrace())))
		}
	}

	t.logger.Error(msg, toZapFieldSlice(args)...)

	// Capture error
	go func(ctx context.Context, err error, fullErrMessage string) {
		captureException(ctx, err, msg)
	}(ctx, originErr, err.Error())
}

func (t tracer) With(args ...Field) Tracer {
	return tracer{
		logger:      t.logger.With(toZapFieldSlice(args)...),
		extraFields: append(t.extraFields, args...),
	}
}

func captureException(ctx context.Context, err error, msg string) {
	hub := sentry.GetHubFromContext(ctx)
	if hub != nil {
		client, scope := hub.Client(), hub.Scope()
		if client == nil || scope == nil {
			return // Skip if client not provided
		}

		event := client.EventFromException(err, sentry.LevelError)
		event.Message = msg
		client.CaptureEvent(event, &sentry.EventHint{
			Context: ctx,
		}, scope)
	}
}
