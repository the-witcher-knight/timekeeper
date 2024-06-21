package auth

import (
	"context"
)

const (
	userProfileCtxKey = "user_profile"
)

func SetInCtx(ctx context.Context, userProfile UserProfile) context.Context {
	return context.WithValue(ctx, userProfileCtxKey, userProfile)
}

func FromCtx(ctx context.Context) UserProfile {
	u, exists := ctx.Value(userProfileCtxKey).(UserProfile)
	if !exists {
		return UserProfile{}
	}

	return u
}
