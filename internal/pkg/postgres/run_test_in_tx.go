package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
)

// LoadSqlTestFile read the sql test file and exec it
func LoadSqlTestFile(t *testing.T, tx ContextExecutor, sqlFile string) {
	b, err := os.ReadFile(sqlFile)
	require.NoError(t, err)

	_, err = tx.Exec(string(b))
	require.NoError(t, err)
}

// RunTestInTx run a test function within a transaction
func RunTestInTx(t *testing.T, testFn func(t *testing.T, tx ContextExecutor)) {
	cfg := config.DBConfig{
		URL:          os.Getenv("DB_URL"),
		MaxOpenConns: 50,
	}
	db, err := Connect(cfg)
	require.NoError(t, err)
	defer db.Close()

	require.NoError(t, runTestInTx(t, db, testFn))
}

func runTestInTx(t *testing.T, db ContextBeginner, testFn func(t *testing.T, tx ContextExecutor)) error {
	// Begin transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	// Rollback when finish
	defer tx.Rollback()

	// Execute the function
	testFn(t, tx)

	return nil
}
