package model

import (
	"time"
)

// Attendance represents model for attendance
type Attendance struct {
	ID          int64
	EmployerID  int64
	CheckInTime time.Time
	Notes       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
