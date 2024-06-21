package httpio

import (
	"net/http"
)

var (
	ErrInternalServerError = Error{Status: http.StatusInternalServerError, Code: "internal_server_error", Desc: "internal server error"}
)
