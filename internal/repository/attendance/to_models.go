package attendance

import (
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository/ormmodel"
)

func toAttendanceModel(o ormmodel.Attendance) model.Attendance {
	return model.Attendance{
		ID:          o.ID,
		EmployerID:  o.EmployerID,
		CheckInTime: o.CheckInTime,
		Notes:       o.Notes,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
		DeletedAt:   o.DeletedAt.Ptr(),
	}
}

func toAttendanceSlice(orms ormmodel.AttendanceSlice) []model.Attendance {
	rs := make([]model.Attendance, len(orms))
	for idx, orm := range orms {
		rs[idx] = toAttendanceModel(*orm)
	}

	return rs
}
