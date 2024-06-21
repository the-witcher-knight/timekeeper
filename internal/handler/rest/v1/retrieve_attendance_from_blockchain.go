package v1

import (
	"net/http"
	"time"

	"github.com/the-witcher-knight/timekeeper/internal/controller/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) RetrieveAttendanceFromBlockchain() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var req getAttendanceRequest
		if err := req.FromURLQuery(r); err != nil {
			return toAPIError(err)
		}

		atts, err := hdl.attCtrl.RetrieveAttendanceFromBlockchain(ctx, attendance.FilterInput{
			EmployerID: req.EmployerID,
			FromTime:   req.FromTime,
			ToTime:     req.ToTime,
		})
		if err != nil {
			return toCtrlError(err)
		}

		respond.OK(toAttendanceFromBlockchainResponseSlice(atts)).Write(w, r)
		return nil
	})
}

type attendanceBlockchainResponse struct {
	ID          int64     `json:"id"`
	EmployerID  int64     `json:"employer_id"`
	CheckInTime time.Time `json:"check_in_time"`
	Notes       string    `json:"notes"`
}

func toAttendanceBlockchainResponse(att model.Attendance) attendanceBlockchainResponse {
	return attendanceBlockchainResponse{
		ID:          att.ID,
		EmployerID:  att.EmployerID,
		CheckInTime: att.CheckInTime,
		Notes:       att.Notes,
	}
}

func toAttendanceFromBlockchainResponseSlice(atts []model.Attendance) []attendanceBlockchainResponse {
	rs := make([]attendanceBlockchainResponse, len(atts))
	for i := range atts {
		rs[i] = toAttendanceBlockchainResponse(atts[i])
	}

	return rs
}
