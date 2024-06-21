package tracing

import (
	"github.com/the-witcher-knight/timekeeper/internal/pkg/logging"
)

func NewNoop() Tracer {
	return tracer{
		logger: logging.Noop(),
	}
}
