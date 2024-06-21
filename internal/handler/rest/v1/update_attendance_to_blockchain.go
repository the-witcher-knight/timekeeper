package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) UpdateAttendanceToBlockchain() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var attID int64
		if idStr := chi.URLParam(r, "id"); len(idStr) > 0 {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return toAPIError(err)
			}

			attID = id
		}

		// Decode the request body into ClientRequest struct
		req, err := httpio.DecodeValidate[updateAttendanceRequest](w, r)
		if err != nil {
			return toAPIError(err)
		}

		if req.ID != attID {
			return errInvalidAttendanceID
		}

		ctrlInput := model.Attendance{
			ID:          req.ID,
			EmployerID:  req.EmployerID,
			CheckInTime: req.CheckInTime,
			Notes:       req.Notes,
		}

		if err := hdl.attCtrl.UpdateAttendanceToBlockchain(ctx, ctrlInput); err != nil {
			return toCtrlError(err)
		}

		respond.OK(httpio.Message{
			Code: "updated",
			Desc: "Attendance updated successfully",
		}).Write(w, r)
		return nil
	})
}

type updateAttendanceRequest struct {
	ID          int64     `json:"id"`
	EmployerID  int64     `json:"employer_id"`
	CheckInTime time.Time `json:"check_in_time"`
	Notes       string    `json:"notes"`
}

func (req updateAttendanceRequest) Valid(ctx context.Context) error {
	if req.ID == 0 {
		return errInvalidAttendanceID
	}

	if req.EmployerID == 0 {
		return errInvalidEmployerID
	}

	return nil
}
