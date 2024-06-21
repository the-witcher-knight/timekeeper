package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	v1 "github.com/the-witcher-knight/timekeeper/internal/handler/rest/v1"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/auth"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/logging"
)

func addRoutes(
	cfg config.AppConfig,
	logger logging.Logger,
	mux *chi.Mux,
	handlerV1 v1.Handler,
) {
	mux.Use(httpio.RootMiddleware(cfg, logger))
	mux.Route("/api", func(r chi.Router) {
		r.Route("/_", func(blankRouter chi.Router) {
			blankRouter.Get("/panic", panicRoute())
			blankRouter.Get("/ping", pingRoute())
		})

		r.Mount("/v1", v1Route(cfg, handlerV1))
		r.Mount("/helpers", helperRoute(handlerV1))
	})
}

func panicRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("simulated panic error")
	}
}

func pingRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = fmt.Fprintln(w, "ok")
	}
}

func v1Route(cfg config.AppConfig, hdl v1.Handler) http.Handler {
	return chi.NewRouter().Route("/v1", func(v1 chi.Router) {
		v1.Mount("/blockchain", blockchainRoute(cfg, hdl))
		v1.Mount("/attendances", attendanceRoute(cfg, hdl))
		v1.Mount("/auth", authRoute(hdl))
	})
}

func blockchainRoute(cfg config.AppConfig, hdl v1.Handler) http.Handler {
	return chi.NewRouter().Route("/blockchain", func(bc chi.Router) {
		bc.Group(func(r chi.Router) {
			r.Use(auth.Middleware(cfg))
			r.Get("/attendances", hdl.RetrieveAttendanceFromBlockchain())
			r.Post("/attendances", hdl.RecordAttendanceToBlockchain())
		})

		bc.Group(func(r chi.Router) {
			r.Use(auth.Middleware(cfg, auth.HasRole(string(model.UserRoleAdmin))))

			r.Put("/attendances/{id}", hdl.UpdateAttendanceToBlockchain())
			r.Post("/accounts/authorize", hdl.AuthorizeAccount())
			r.Delete("/accounts/deauthorize", hdl.DeauthorizeAccount())
		})
	})
}

func attendanceRoute(cfg config.AppConfig, hdl v1.Handler) http.Handler {
	return chi.NewRouter().Route("/attendances", func(bc chi.Router) {
		bc.Use(auth.Middleware(cfg))
		bc.Get("/", hdl.GetAttendances())
	})
}

func authRoute(hdl v1.Handler) http.Handler {
	return chi.NewRouter().Route("/auth", func(a chi.Router) {
		a.Post("/token", hdl.SignIn())
	})
}

func helperRoute(hdl v1.Handler) http.Handler {
	return chi.NewRouter().Route("/helpers", func(hp chi.Router) {
		hp.Post("/users/employee", hdl.CreateEmployee())
		hp.Post("/users/admin", hdl.CreateAdmin())
	})
}
