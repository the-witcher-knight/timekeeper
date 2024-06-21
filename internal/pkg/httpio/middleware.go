package httpio

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/logging"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/tracing"
)

func RootMiddleware(cfg config.AppConfig, logger logging.Logger) func(next http.Handler) http.Handler {
	tracer := tracing.NewTracer(logger)

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// Add sentry to context
			ctx = sentry.SetHubOnContext(ctx, sentry.CurrentHub())

			defer func() {
				// Recover from the panic and obtain the panic value.
				if p := recover(); p != nil {
					err, ok := p.(error)
					if !ok {
						err = fmt.Errorf("%+v", p)
					}

					// Capture and log the entire stack trace along with the error details.
					tracer.Error(ctx, err, "caught a panic", false,
						tracing.String("stacktrace", string(debug.Stack())),
					)

					// Respond with a 500 Internal Server Error and log any encoding errors.
					WriteJSON(w, r, Response[Error]{
						Status: http.StatusInternalServerError,
						Body:   ErrInternalServerError,
					})
				}
			}()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r.WithContext(tracing.SetInCtx(ctx, tracer)))

			tracer.With(
				tracing.String("host.name", r.Host),
				tracing.String("url.path", r.URL.Path),
				tracing.String("url.query", r.URL.RawQuery),
				tracing.String("http.request.method_original", r.Method),
				tracing.Int("http.request.body.size", int(r.ContentLength)),
				tracing.String("http.request.proto", r.Proto),
				tracing.String("http.request.remote_address", r.RemoteAddr),
				tracing.String("user_agent.original", r.UserAgent()),
				tracing.Int("http.response.status_code", ww.Status()),
				tracing.Int("http.response.body.size", ww.BytesWritten()),
			).Info(ctx, "Served")
		}

		return http.HandlerFunc(fn)
	}
}
