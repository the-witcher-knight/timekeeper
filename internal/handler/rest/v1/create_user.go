package v1

import (
	"context"
	"net/http"

	"github.com/the-witcher-knight/timekeeper/internal/controller/user"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) CreateEmployee() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		req, err := httpio.DecodeValidate[createUserRequest](w, r)
		if err != nil {
			return toAPIError(err)
		}

		ctrlInput := user.InputUser{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}
		if err := hdl.userCtrl.CreateEmployee(ctx, ctrlInput); err != nil {
			return toCtrlError(err)
		}

		respond.OK(httpio.Message{
			Code: "user_created",
			Desc: "Employee created successfully",
		}).Write(w, r)
		return nil
	})
}

func (hdl Handler) CreateAdmin() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		req, err := httpio.DecodeValidate[createUserRequest](w, r)
		if err != nil {
			return toAPIError(err)
		}

		ctrlInput := user.InputUser{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}
		if err := hdl.userCtrl.CreateAdmin(ctx, ctrlInput); err != nil {
			return toCtrlError(err)
		}

		respond.OK(httpio.Message{
			Code: "user_created",
			Desc: "Admin created successfully",
		}).Write(w, r)
		return nil
	})
}

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req createUserRequest) Valid(ctx context.Context) error {
	if len(req.Name) == 0 {
		return errUserNameCannotBeBlank
	}

	if len(req.Email) == 0 {
		return errUserEmailCannotBeBlank
	}

	if len(req.Password) == 0 {
		return errUserPasswordCannotBeBlank
	}

	return nil
}
