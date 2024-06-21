package logging

import (
	"go.uber.org/zap"
)

type Option func(cfg *zap.Config)

func WithOutputPaths(paths ...string) Option {
	return func(cfg *zap.Config) {
		cfg.OutputPaths = append(cfg.OutputPaths, paths...)
	}
}

func WithErrorOutputPaths(paths ...string) Option {
	return func(cfg *zap.Config) {
		cfg.OutputPaths = append(cfg.OutputPaths, paths...)
	}
}

func WithExtraFields(fieldMap map[string]interface{}) Option {
	return func(cfg *zap.Config) {
		cfg.InitialFields = fieldMap
	}
}
