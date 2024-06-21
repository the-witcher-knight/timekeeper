package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"

	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/logging"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/waiter"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/postgres"
)

// System represents the core components of an application, including its configuration,
// database connection, HTTP request router, and a waiter for graceful shutdown. It serves
// as the central struct that encapsulates these essential elements for managing and
// controlling the application's behavior.
type System struct {
	cfg           config.AppConfig
	db            *sql.DB
	mux           *chi.Mux
	ethClientHTTP *ethclient.Client
	ethClientWS   *ethclient.Client
	logger        logging.Logger
	waiter        waiter.Waiter
}

func New(cfg config.AppConfig) (*System, error) {
	s := &System{cfg: cfg}

	if err := s.initDB(); err != nil {
		return nil, err
	}

	if err := s.pingDB(); err != nil {
		return nil, pkgerrors.Wrapf(err, "failed to ping DB")
	}

	s.initMux()

	if err := s.initEthClientHTTP(); err != nil {
		return nil, pkgerrors.Wrapf(err, "failed to init eth client http")
	}

	if err := s.initEthClientWS(); err != nil {
		return nil, pkgerrors.Wrapf(err, "failed to init eth client ws")
	}

	if err := s.initLogger(); err != nil {
		return nil, pkgerrors.Wrapf(err, "failed to init logger")
	}

	if err := s.initSentry(); err != nil {
		return nil, pkgerrors.Wrapf(err, "failed to init sentry")
	}

	s.initWaiter()
	s.waiter.Cleanup(func() {
		// Close db connection
		s.logger.Info("close db connection")
		if err := s.db.Close(); err != nil {
			s.logger.Error(fmt.Sprintf("close db connection error: %v", err))
		}

		// Flush sentry
		sentry.Flush(time.Second * 5)

		// Flush all log
		_ = s.logger.Flush()
	})

	return s, nil
}

func (s *System) Config() config.AppConfig {
	return s.cfg
}

func (s *System) initDB() (err error) {
	s.db, err = postgres.Connect(s.Config().DB)

	return
}

func (s *System) pingDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.db.PingContext(ctx)
}

func (s *System) DB() postgres.ContextBeginner {
	return s.db
}

func (s *System) initMux() {
	s.mux = chi.NewMux()
}

func (s *System) Mux() *chi.Mux {
	return s.mux
}

func (s *System) initLogger() (err error) {
	s.logger, err = logging.New(s.Config())

	return
}

func (s *System) Logger() logging.Logger {
	return s.logger
}

func (s *System) initEthClientHTTP() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := ethclient.DialContext(ctx, s.Config().BlockChain.NetworkHTTP)
	if err != nil {
		return err
	}

	s.ethClientHTTP = conn
	return nil
}

func (s *System) EthClientHTTP() *ethclient.Client {
	return s.ethClientHTTP
}

func (s *System) initEthClientWS() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := ethclient.DialContext(ctx, s.Config().BlockChain.NetworkWS)
	if err != nil {
		return err
	}

	s.ethClientHTTP = conn
	return nil
}

func (s *System) EthClientWS() *ethclient.Client {
	return s.ethClientHTTP
}

func (s *System) initWaiter() {
	s.waiter = waiter.New(waiter.CatchSignals())
}

func (s *System) Waiter() waiter.Waiter {
	return s.waiter
}

func (s *System) initSentry() error {
	sentryDSN := s.Config().Sentry.DSN
	if len(sentryDSN) == 0 {
		return nil
	}

	return sentry.Init(sentry.ClientOptions{
		Dsn:           s.Config().Sentry.DSN,
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
}

func (s *System) WaitForWeb(ctx context.Context) error {
	webServer := &http.Server{
		Addr:    s.cfg.Web.Address(),
		Handler: s.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Printf("web server started; listening at http://localhost%s\n", s.cfg.Web.Port)
		defer fmt.Println("web server shutdown")
		if err := webServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return group.Wait()
}
