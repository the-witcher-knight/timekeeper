package main

import (
	"fmt"
	"os"

	"github.com/the-witcher-knight/timekeeper/cmd/asciiart"
	"github.com/the-witcher-knight/timekeeper/internal/controller/user"
	v1 "github.com/the-witcher-knight/timekeeper/internal/handler/rest/v1"
	"github.com/the-witcher-knight/timekeeper/internal/ids"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/system"
	"github.com/the-witcher-knight/timekeeper/internal/repository"
)

func main() {
	asciiart.Show()
	if err := run(); err != nil {
		fmt.Printf("server exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.ReadConfigFromEnv()
	if err != nil {
		return err
	}

	s, err := system.New(cfg)
	if err != nil {
		return err
	}

	if err := ids.Setup(); err != nil {
		return err
	}

	store := repository.New(s.DB())
	userCtrl := user.New(store, []byte(cfg.JWTSecret))
	handlerV1 := v1.New(userCtrl)

	addRoutes(
		s.Config(),
		s.Logger(),
		s.Mux(),
		handlerV1,
	)

	s.Waiter().Add(
		s.WaitForWeb,
		// Add more waiter func, e.g. WaitForGRPC, WaitForJWKsPolling,...
	)

	fmt.Printf("started %s application\n", cfg.ProjectName)
	defer fmt.Printf("stopped %s application\n", cfg.ProjectName)

	return s.Waiter().Wait()
}
