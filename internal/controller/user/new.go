package user

import (
	"context"

	"github.com/the-witcher-knight/timekeeper/internal/repository"
)

type Controller interface {
	// CreateEmployee creates an employee user
	CreateEmployee(ctx context.Context, user InputUser) error

	// CreateAdmin creates an admin user
	CreateAdmin(ctx context.Context, user InputUser) error

	// SignIn generate access token
	SignIn(ctx context.Context, user SignInInput) (string, error)
}

type controller struct {
	repo      repository.Registry
	jwtSecret []byte
}

func New(repo repository.Registry, jwtSecret []byte) Controller {
	return &controller{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}
