package user

import (
	"context"
	"strconv"
	"time"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/auth"
	"github.com/the-witcher-knight/timekeeper/internal/repository/user"
)

var (
	timeNowWrapper = func() time.Time {
		return time.Now().UTC()
	}
)

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ctrl controller) SignIn(ctx context.Context, input SignInInput) (string, error) {
	// Find user by email
	existed, err := ctrl.repo.User().GetUserByFilter(ctx, user.FilterInput{
		Email: input.Email,
	}, false)
	if err != nil {
		return "", err
	}

	if existed.ID == 0 {
		return "", ErrUserEmailNotExists
	}

	// Compare with hashed password
	match, err := comparePasswordsWrapper(existed.Password, input.Password)
	if err != nil {
		return "", err
	}

	// Return error if password not match
	if !match {
		return "", ErrPasswordNotMatch
	}

	now := timeNowWrapper()
	claims := auth.Claims{
		Sub: strconv.Itoa(int(existed.ID)),
		Exp: now.Add(time.Hour).Unix(),
		Iat: now.Unix(),
		ExtraClaims: map[string]interface{}{
			"id":   existed.ID,
			"role": existed.Role,
		},
	}
	// Generate access token
	token, err := auth.GenerateToken(ctrl.jwtSecret, claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
