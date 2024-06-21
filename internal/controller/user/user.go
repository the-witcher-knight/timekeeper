package user

import (
	"context"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository/user"
)

// InputUser contains name and email of a user
type InputUser struct {
	Name     string
	Email    string
	Password string
}

func (ctrl controller) CreateEmployee(ctx context.Context, input InputUser) error {
	roles := []model.UserRole{
		model.UserRoleEmployee,
	}

	return createUser(ctx, ctrl, input, roles)
}

func (ctrl controller) CreateAdmin(ctx context.Context, input InputUser) error {
	roles := []model.UserRole{
		model.UserRoleEmployee, // Include employee role
		model.UserRoleAdmin,
	}

	return createUser(ctx, ctrl, input, roles)
}

// CreateUser creates an employee user by given input
func createUser(ctx context.Context, ctrl controller, input InputUser, roles []model.UserRole) error {
	// Find user by email
	existed, err := ctrl.repo.User().GetUserByFilter(ctx, user.FilterInput{
		Email: input.Email,
	}, false)
	if err != nil {
		return err
	}

	if existed.ID != 0 {
		return ErrEmailExisted
	}

	// Hash given password
	pwd, err := hashPasswordWrapper(input.Password)
	if err != nil {
		return err
	}

	if err := ctrl.repo.User().CreateUser(ctx, model.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: pwd,
		Role:     roles,
	}); err != nil {
		return err
	}

	return nil
}
