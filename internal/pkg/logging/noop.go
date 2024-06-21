package logging

import (
	"go.uber.org/zap"
)

func Noop() Logger {
	return Logger{
		zap: zap.NewNop(),
	}
}
