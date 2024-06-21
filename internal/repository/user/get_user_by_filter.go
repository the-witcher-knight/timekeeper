package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository/ormmodel"
)

type FilterInput struct {
	Name  string
	Email string
}

func (repo repo) GetUserByFilter(ctx context.Context, filters FilterInput, lock bool) (model.User, error) {
	qms := []qm.QueryMod{
		ormmodel.UserWhere.DeletedAt.IsNull(),
	}

	if filters.Name != "" {
		qms = append(qms, ormmodel.UserWhere.Name.LIKE("%"+filters.Name+"%"))
	}

	if filters.Email != "" {
		qms = append(qms, ormmodel.UserWhere.Email.EQ(filters.Email))
	}

	if lock {
		qms = append(qms, qm.For("UPDATE"))
	}

	o, err := ormmodel.Users(qms...).One(ctx, repo.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, nil
		}
		return model.User{}, pkgerrors.WithStack(err)
	}

	return toUserModel(*o), nil
}
