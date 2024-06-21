package user

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/the-witcher-knight/timekeeper/internal/ids"
	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository/ormmodel"
)

func (repo repo) CreateUser(ctx context.Context, user model.User) error {
	if user.ID == 0 {
		userID, err := ids.User.NextID()
		if err != nil {
			return pkgerrors.WithStack(err)
		}

		user.ID = userID
	}

	roles := make([]string, len(user.Role))
	for idx, role := range user.Role {
		roles[idx] = string(role)
	}

	orm := ormmodel.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      roles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: null.TimeFromPtr(user.DeletedAt),
	}
	if err := orm.Insert(ctx, repo.dbConn, boil.Infer()); err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
