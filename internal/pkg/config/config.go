package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/constants"
)

// AppConfig representing an application configuration
type AppConfig struct {
	ProjectName     string
	Environment     string
	DB              DBConfig
	Web             WebConfig
	Sentry          SentryConfig
	BlockChain      BlockChainConfig
	JWTSecret       string
	ShutdownTimeout time.Duration
}

// DBConfig representing a postgres configuration
type DBConfig struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
}

// WebConfig representing a web configuration
type WebConfig struct {
	Host string
	Port string
}

func (c WebConfig) Address() string {
	// Make sure the port has the format :8080
	return fmt.Sprintf("%s%v", c.Host, c.Port)
}

// SentryConfig representing a sentry configuration
type SentryConfig struct {
	DSN string
}

// BlockChainConfig representing a blockchain configuration
type BlockChainConfig struct {
	NetworkHTTP     string
	NetworkWS       string
	PrivateKey      string
	ContractAddress string
}

// ReadConfigFromEnv reads all environment variables, validates it and parses it into AppConfig struct
func ReadConfigFromEnv() (AppConfig, error) {
	projectName := os.Getenv("PROJECT_NAME")
	if projectName == "" {
		return AppConfig{}, ErrInvalidProjectName
	}

	environment := os.Getenv("ENVIRONMENT")
	if err := validateEnvironment(environment); err != nil {
		return AppConfig{}, ErrInvalidEnvironment
	}

	port, err := strconv.Atoi(strings.TrimSpace(os.Getenv("APP_PORT")))
	if err != nil || (port < 0 || port > 99999) {
		return AppConfig{}, ErrInvalidPort
	}

	dbURL := strings.TrimSpace(os.Getenv("DB_URL"))
	if dbURL == "" {
		return AppConfig{}, ErrInvalidDBURL
	}

	var maxOpenConns int
	if maxOpenConnsStr := strings.TrimSpace(os.Getenv("DB_MAX_OPEN_CONN")); maxOpenConnsStr != "" {
		maxOpenConns, err = strconv.Atoi(maxOpenConnsStr)
		if err != nil {
			return AppConfig{}, pkgerrors.Wrap(err, "invalid DB_MAX_OPEN_CONN")
		}
	}

	var maxIdleConns int
	if maxIdleConnsStr := strings.TrimSpace(os.Getenv("DB_MAX_IDLE_CONN")); maxIdleConnsStr != "" {
		maxIdleConns, err = strconv.Atoi(maxIdleConnsStr)
		if err != nil {
			return AppConfig{}, pkgerrors.Wrap(err, "invalid DB_MAX_IDLE_CONN")
		}
	}

	// Sentry DSN can be empty
	sentryDSN := strings.TrimSpace(os.Getenv("SENTRY_DSN"))
	if sentryDSN == "" {
		fmt.Println("sentry dsn dnv not set. the app will not capture exceptions to sentry")
	}

	networkHTTP := strings.TrimSpace(os.Getenv("BLOCKCHAIN_NETWORK_HTTP"))
	if networkHTTP == "" {
		return AppConfig{}, ErrInvalidBlockChainNetwork
	}

	networkWS := strings.TrimSpace(os.Getenv("BLOCKCHAIN_NETWORK_WS"))
	if networkWS == "" {
		return AppConfig{}, ErrInvalidBlockChainNetwork
	}

	privateKey := strings.TrimSpace(os.Getenv("BLOCKCHAIN_PRIVATE_KEY"))
	if privateKey == "" {
		return AppConfig{}, ErrInvalidBlockChainPrivateKey
	}

	contractAddress := strings.TrimSpace(os.Getenv("BLOCKCHAIN_CONTRACT_ADDRESS"))
	if contractAddress == "" {
		// It may empty before deployment
		fmt.Printf("%s on API service\n", ErrInvalidBlockChainContractAddress)
	}

	jwtSecret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if jwtSecret == "" {
		return AppConfig{}, ErrInvalidJWTSecret
	}

	return AppConfig{
		ProjectName: projectName,
		Environment: environment,
		Web: WebConfig{
			Host: "0.0.0.0",
			Port: fmt.Sprintf(":%v", port),
		},
		DB: DBConfig{
			URL:          dbURL,
			MaxOpenConns: maxOpenConns,
			MaxIdleConns: maxIdleConns,
		},
		Sentry: SentryConfig{
			DSN: sentryDSN,
		},
		BlockChain: BlockChainConfig{
			NetworkHTTP:     networkHTTP,
			NetworkWS:       networkWS,
			PrivateKey:      privateKey,
			ContractAddress: contractAddress,
		},
		JWTSecret: jwtSecret,
	}, nil
}

func validateEnvironment(str string) error {
	envs := []constants.Environment{constants.EnvironmentDevelopment, constants.EnvironmentProduction}
	for _, env := range envs {
		if str == string(env) {
			return nil
		}
	}

	return ErrInvalidEnvironment
}
