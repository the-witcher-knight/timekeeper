package user

import (
	"context"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/postgres"
)

type Repository interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user model.User) error

	// GetUserByFilter gets a user by given criteria
	GetUserByFilter(ctx context.Context, filters FilterInput, lock bool) (model.User, error)
}

type repo struct {
	dbConn postgres.ContextExecutor
}

func New(db postgres.ContextExecutor) Repository {
	return &repo{
		dbConn: db,
	}
}
