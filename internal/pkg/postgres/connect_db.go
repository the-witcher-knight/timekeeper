package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
)

// Connect returns the singleton instance of the database
func Connect(cfg config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.URL)
	if err != nil {
		return nil, err
	}

	// Set default value for max idle connections
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 1
	}

	// Set default value for max open connections
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 10
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}
