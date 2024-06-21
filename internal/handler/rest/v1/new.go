package v1

import (
	"github.com/the-witcher-knight/timekeeper/internal/controller/user"
)

type Handler struct {
	userCtrl user.Controller
}

func New(userCtrl user.Controller) Handler {
	return Handler{
		userCtrl: userCtrl,
	}
}
