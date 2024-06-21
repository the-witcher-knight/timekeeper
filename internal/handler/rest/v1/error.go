package v1

import (
	"errors"
	"net/http"

	"github.com/the-witcher-knight/timekeeper/internal/controller/user"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
)

var (
	errEmailExisted              = httpio.Error{Status: http.StatusBadRequest, Code: "email_existed", Desc: "email is already exists"}
	errUserNameCannotBeBlank     = httpio.Error{Status: http.StatusBadRequest, Code: "invalid_request", Desc: "user name can not be blank"}
	errUserEmailCannotBeBlank    = httpio.Error{Status: http.StatusBadRequest, Code: "invalid_request", Desc: "user email	can not be blank"}
	errUserPasswordCannotBeBlank = httpio.Error{Status: http.StatusBadRequest, Code: "invalid_request", Desc: "user password can not be blank"}
	errEmailOrPasswordNotValid   = httpio.Error{Status: http.StatusBadRequest, Code: "invalid_request", Desc: "email or password not valid"}
)

func toAPIError(err error) error {
	var apiErr httpio.Error
	if errors.As(err, &apiErr) {
		return apiErr
	}

	return httpio.Error{Status: http.StatusBadRequest, Code: "invalid_request", Desc: err.Error()}
}

func toCtrlError(err error) error {
	switch err.Error() {
	case user.ErrEmailExisted.Error():
		return errEmailExisted
	case user.ErrUserEmailNotExists.Error(), user.ErrPasswordNotMatch.Error():
		return errEmailOrPasswordNotValid

	default:
		return err // This error will be handled in HandlerFunc func
	}
}
