package attendance

import (
	"context"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository"
)

type Controller interface {
	// RecordAttendanceToBlockchain records an attendance to blockchain by given employer ID
	RecordAttendanceToBlockchain(ctx context.Context, employerID int64) error

	// UpdateAttendanceToBlockchain update an attendance to blockchain by given attendance information
	UpdateAttendanceToBlockchain(ctx context.Context, att model.Attendance) error

	// RetrieveAttendanceFromBlockchain returns a list of attendance from blockchain by given filter
	RetrieveAttendanceFromBlockchain(ctx context.Context, filter FilterInput) ([]model.Attendance, error)

	// CreateAttendance creates a new attendance to database
	CreateAttendance(ctx context.Context, att model.Attendance) error

	// ListAttendances returns a list of attendance from database by given filter
	ListAttendances(ctx context.Context, filter FilterInput) ([]model.Attendance, error)
}

type controller struct {
	bc   blockchain.Blockchain
	repo repository.Registry
}

func New(bc blockchain.Blockchain, repo repository.Registry) Controller {
	return &controller{
		bc:   bc,
		repo: repo,
	}
}
