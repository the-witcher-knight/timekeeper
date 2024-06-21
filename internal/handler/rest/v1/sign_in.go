package v1

import (
	"context"
	"net/http"

	"github.com/the-witcher-knight/timekeeper/internal/controller/user"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) SignIn() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		req, err := httpio.DecodeValidate[signInRequest](w, r)
		if err != nil {
			return toAPIError(err)
		}

		ctrlInput := user.SignInInput{
			Email:    req.Email,
			Password: req.Password,
		}
		token, err := hdl.userCtrl.SignIn(ctx, ctrlInput)
		if err != nil {
			return toCtrlError(err)
		}

		respond.OK(signInResponse{
			Token: token,
		}).Write(w, r)
		return nil
	})
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req signInRequest) Valid(ctx context.Context) error {
	if len(req.Email) == 0 {
		return errUserEmailCannotBeBlank
	}

	if len(req.Password) == 0 {
		return errUserPasswordCannotBeBlank
	}

	return nil
}

type signInResponse struct {
	Token string `json:"token"`
}
