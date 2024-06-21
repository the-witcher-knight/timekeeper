package logging

import (
	"errors"
	"syscall"

	"go.uber.org/zap"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/constants"
)

type Logger struct {
	zap *zap.Logger
}

func New(cfg config.AppConfig, opts ...Option) (Logger, error) {
	zapConfig := newZapConfig("stdout", "stderr")

	if cfg.Environment == string(constants.EnvironmentProduction) {
		zapConfig.Development = false
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	for _, opt := range opts {
		opt(&zapConfig)
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		return Logger{}, err
	}

	return Logger{
		zap: zapLogger,
	}, nil
}

func (l Logger) Debug(msg string, args ...zap.Field) {
	l.zap.Debug(msg, args...)
}

func (l Logger) Error(msg string, args ...zap.Field) {
	l.zap.Error(msg, args...)
}

func (l Logger) Info(msg string, args ...zap.Field) {
	l.zap.Info(msg, args...)
}

func (l Logger) With(args ...zap.Field) Logger {
	return Logger{
		zap: l.zap.With(args...),
	}
}

func (l Logger) Flush() error {
	if err := l.zap.Sync(); err != nil {
		// Ignore this stderr https://github.com/uber-go/zap/issues/328
		if !errors.Is(err, syscall.ENOTTY) && !errors.Is(err, syscall.EINVAL) {
			return err
		}
	}

	return nil
}
