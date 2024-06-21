package v1

import (
	"github.com/the-witcher-knight/timekeeper/internal/controller/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/controller/bcauth"
	"github.com/the-witcher-knight/timekeeper/internal/controller/user"
)

type Handler struct {
	userCtrl   user.Controller
	attCtrl    attendance.Controller
	bcAuthCtrl bcauth.Controller
}

func New(userCtrl user.Controller, attCtrl attendance.Controller, bcAuth bcauth.Controller) Handler {
	return Handler{
		userCtrl:   userCtrl,
		attCtrl:    attCtrl,
		bcAuthCtrl: bcAuth,
	}
}
