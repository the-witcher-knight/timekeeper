package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapConfig(outputPath, errOutputPath string) zap.Config {
	cfg := zap.Config{
		Development: true,
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:          "json",
		EncoderConfig:     newEncoderConfig(),
		OutputPaths:       []string{outputPath},
		ErrorOutputPaths:  []string{errOutputPath},
		DisableStacktrace: true,
		DisableCaller:     false,
	}

	return cfg
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
