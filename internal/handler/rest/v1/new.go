package v1

import (
	"github.com/the-witcher-knight/timekeeper/internal/controller/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/controller/user"
)

type Handler struct {
	userCtrl user.Controller
	attCtrl  attendance.Controller
}

func New(userCtrl user.Controller, attCtrl attendance.Controller) Handler {
	return Handler{
		userCtrl: userCtrl,
		attCtrl:  attCtrl,
	}
}
