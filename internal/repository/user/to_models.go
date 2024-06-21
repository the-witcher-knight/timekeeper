package user

import (
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository/ormmodel"
)

func toUserModel(orm ormmodel.User) model.User {
	return model.User{
		ID:        orm.ID,
		Name:      orm.Name,
		Email:     orm.Email,
		Password:  orm.Password,
		Role:      toUserRoleSlice(orm.Role),
		CreatedAt: orm.CreatedAt,
		UpdatedAt: orm.UpdatedAt,
		DeletedAt: orm.DeletedAt.Ptr(),
	}
}

func toUserModelSlice(orms ormmodel.UserSlice) []model.User {
	rs := make([]model.User, len(orms))
	for i := range orms {
		rs[i] = toUserModel(*orms[i])
	}

	return rs
}

func toUserRoleModel(role string) model.UserRole {
	return model.UserRole(role)
}

func toUserRoleSlice(roles []string) []model.UserRole {
	rs := make([]model.UserRole, len(roles))
	for i := range roles {
		rs[i] = toUserRoleModel(roles[i])
	}

	return rs
}
