package attendance

import (
	"context"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/postgres"
)

type Repository interface {
	// CreateAttendance creates a new attendance
	CreateAttendance(ctx context.Context, att model.Attendance) error

	// GetAttendanceByID gets an attendance by given id
	GetAttendanceByID(ctx context.Context, id int64, lock bool) (model.Attendance, error)

	// ListAttendanceByFilter gets a user by given criteria
	ListAttendanceByFilter(ctx context.Context, filters FilterInput) ([]model.Attendance, error)

	// UpdateAttendance updates a existed attendance
	UpdateAttendance(ctx context.Context, att model.Attendance) error
}

type repository struct {
	dbConn postgres.ContextExecutor
}

func New(db postgres.ContextExecutor) Repository {
	return &repository{
		dbConn: db,
	}
}
