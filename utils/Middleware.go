package utils

import (
	"context"
	"fmt"
	"myapp/graph/model"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func GetUserByUsername(username string) (*model.User, error) {
	var user *model.User

	user = &model.User{
		Email:    "louisaldorio@gmail.com",
		Role:     "distributor",
		Username: "louisaldorio",
	}

	return user, nil
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")


			// Allow unauthenticated users in
			if authToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			jwtToken, err := ValidateToken(authToken)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			//validate claim
			claims, ok := jwtToken.Claims.(*UserClaim)
			if !ok && !jwtToken.Valid {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			//get user dataclaims
			username := fmt.Sprintf("%v", claims.Username)
			user, err := GetUserByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			//return user data to req
			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			reqWithCtx := r.WithContext(ctx)
			next.ServeHTTP(w, reqWithCtx)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
