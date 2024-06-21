package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/the-witcher-knight/timekeeper/internal/controller/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) GetAttendances() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var req getAttendanceRequest
		if err := req.FromURLQuery(r); err != nil {
			return toAPIError(err)
		}

		atts, err := hdl.attCtrl.ListAttendances(ctx, attendance.FilterInput{
			EmployerID: req.EmployerID,
			FromTime:   req.FromTime,
			ToTime:     req.ToTime,
		})
		if err != nil {
			return toCtrlError(err)
		}

		respond.OK(toAttendanceResponseSlice(atts)).
			Write(w, r)
		return nil
	})
}

type getAttendanceRequest struct {
	EmployerID int64     `json:"employer_id"`
	FromTime   time.Time `json:"from_time"`
	ToTime     time.Time `json:"to_time"`
}

func (req *getAttendanceRequest) FromURLQuery(r *http.Request) error {
	if employerIDStr := r.URL.Query().Get("employer_id"); employerIDStr != "" {
		employerID, err := strconv.ParseInt(employerIDStr, 10, 64)
		if err != nil {
			return err
		}

		req.EmployerID = employerID
	}

	if fromTimeStr := r.URL.Query().Get("from_time"); fromTimeStr != "" {
		fromTime, err := time.Parse(time.RFC3339, fromTimeStr)
		if err != nil {
			return err
		}

		req.FromTime = fromTime
	}

	if toTimeStr := r.URL.Query().Get("to_time"); toTimeStr != "" {
		toTime, err := time.Parse(time.RFC3339, toTimeStr)
		if err != nil {
			return err
		}

		req.ToTime = toTime
	}

	if !(req.FromTime.IsZero() || req.ToTime.IsZero()) &&
		!req.ToTime.After(req.FromTime) {
		return errInvalidTimeRange
	}

	return nil
}

type attendanceResponse struct {
	ID          int64      `json:"id"`
	EmployerID  int64      `json:"employer_id"`
	CheckInTime time.Time  `json:"check_in_time"`
	Notes       string     `json:"notes"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func toAttendanceResponse(m model.Attendance) attendanceResponse {
	return attendanceResponse{
		ID:          m.ID,
		EmployerID:  m.EmployerID,
		CheckInTime: m.CheckInTime,
		Notes:       m.Notes,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}
}

func toAttendanceResponseSlice(ms []model.Attendance) []attendanceResponse {
	rs := make([]attendanceResponse, len(ms))
	for i, m := range ms {
		rs[i] = toAttendanceResponse(m)
	}

	return rs
}
