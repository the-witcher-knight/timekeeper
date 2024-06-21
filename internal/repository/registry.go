package repository

import (
	"context"
	"database/sql"

	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/postgres"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/tracing"
	"github.com/the-witcher-knight/timekeeper/internal/repository/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/repository/user"
)

type Registry interface {
	// User returns the user repository
	User() user.Repository

	// Attendance return the attendance repository
	Attendance() attendance.Repository

	// DoInTx run the func f in a transaction
	DoInTx(ctx context.Context, f func(Registry) error) error
}

type impl struct {
	db         postgres.ContextBeginner
	user       user.Repository
	attendance attendance.Repository
}

func New(db postgres.ContextBeginner) Registry {
	return &impl{
		db:         db,
		user:       user.New(db),
		attendance: attendance.New(db),
	}
}

func (r impl) User() user.Repository {
	return r.user
}

func (r impl) Attendance() attendance.Repository { return r.attendance }

func (r impl) DoInTx(ctx context.Context, f func(repo Registry) error) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	var txerr error
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				tracing.FromCtx(ctx).Error(ctx, err, "rollback when panic error", true)
			}
			panic(p)
		} else if txerr != nil {
			txerr = tx.Rollback()
		} else {
			txerr = tx.Commit()
		}
	}()

	repo := &impl{
		db:         nil, // We don't use registry db here
		user:       user.New(tx),
		attendance: attendance.New(tx),
		// Add more repository here
	}
	txerr = f(repo)

	return txerr
}
