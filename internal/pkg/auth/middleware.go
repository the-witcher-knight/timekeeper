package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
)

func Middleware(cfg config.AppConfig, opts ...OptFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Get token from request header
			token := getTokenFromRequest(r)
			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Verify token and retrieve claims from payload
			claims, err := VerifyToken([]byte(cfg.JWTSecret), token)
			if err != nil {
				httpio.WriteJSON(w, r, httpio.Response[httpio.Message]{
					Status: http.StatusUnauthorized,
					Body: httpio.Message{
						Code: "unauthorized",
						Desc: err.Error(),
					},
				})
				return
			}

			// Parse user profile from claims
			userProfile, err := parseUserProfile(claims)
			if err != nil {
				httpio.WriteJSON(w, r, httpio.Response[httpio.Message]{
					Status: http.StatusUnauthorized,
					Body: httpio.Message{
						Code: "unauthorized",
						Desc: err.Error(),
					},
				})
				return
			}

			// Run option functions
			for _, opt := range opts {
				if err := opt(userProfile); err != nil {
					httpio.WriteJSON(w, r, httpio.Response[httpio.Message]{
						Status: http.StatusUnauthorized,
						Body: httpio.Message{
							Code: "unauthorized",
							Desc: err.Error(),
						},
					})
					return
				}
			}

			// Set user profile to request context
			ctx := SetInCtx(r.Context(), userProfile)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func getTokenFromRequest(r *http.Request) string {
	str := r.Header.Get("Authorization")

	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		return ""
	}

	return parts[1]
}

type UserProfile struct {
	ID    int64
	Email string
	Role  []string
}

func parseUserProfile(claims Claims) (UserProfile, error) {
	idClaim, ok := claims.ExtraClaims["id"].(interface{})
	if !ok {
		return UserProfile{}, ErrInvalidUserProfile
	}

	roleRaws, ok := claims.ExtraClaims["role"].([]interface{})
	if !ok {
		return UserProfile{}, ErrInvalidUserProfile
	}

	roles := make([]string, len(roleRaws))
	for i, role := range roleRaws {
		roles[i] = fmt.Sprintf("%s", role)
	}

	return UserProfile{
		ID:    int64(idClaim.(float64)),
		Email: claims.Sub,
		Role:  roles,
	}, nil
}

type OptFunc func(UserProfile) error

func HasRole(role string) OptFunc {
	return func(userProfile UserProfile) error {
		for _, r := range userProfile.Role {
			if role == r {
				return nil
			}
		}

		return ErrUnauthorized
	}
}
