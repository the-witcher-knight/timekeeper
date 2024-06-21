package blockchain

import (
	"time"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain/contracts"
	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func toAttendanceModel(att contracts.AttendanceAttendanceRecord) model.Attendance {
	return model.Attendance{
		ID:          att.Id.Int64(),
		EmployerID:  att.EmployerId.Int64(),
		CheckInTime: time.Unix(att.CheckInTime.Int64(), 0),
		Notes:       att.Notes,
	}
}

func toAttendanceSlice(att []contracts.AttendanceAttendanceRecord) []model.Attendance {
	result := make([]model.Attendance, len(att))
	for i := range att {
		result[i] = toAttendanceModel(att[i])
	}

	return result
}
