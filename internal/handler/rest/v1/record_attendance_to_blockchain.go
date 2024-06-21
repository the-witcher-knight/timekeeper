package v1

import (
	"net/http"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/auth"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) RecordAttendanceToBlockchain() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		userProfile := auth.FromCtx(ctx)
		if userProfile.ID == 0 {
			return errUserProfileNotFound
		}

		if err := hdl.attCtrl.RecordAttendanceToBlockchain(ctx, userProfile.ID); err != nil {
			return toCtrlError(err)
		}

		respond.OK(httpio.Message{
			Code: "created",
			Desc: "Attendance created successfully",
		}).Write(w, r)
		return nil
	})
}
