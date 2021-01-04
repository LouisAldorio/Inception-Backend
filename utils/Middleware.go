package utils

import (
	"context"
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func GetUserByUsername(username string) (*model.User, error) {
	var user *model.User

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOptions := options.FindOneOptions{}
	cur := collection.FindOne(ctx, bson.D{
		{"username", username},
	}, &findOptions)
	var result bson.M
	err := cur.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	var friendList []string
	for _, v := range result["friendList"].(primitive.A) {
		friend := fmt.Sprintf("%v", v)
		friendList = append(friendList, friend)
	}

	var lookingFor []string
	for _, v := range result["lookingFor"].(primitive.A) {
		item := fmt.Sprintf("%v", v)
		lookingFor = append(lookingFor, item)
	}

	user = &model.User{
		Username:       fmt.Sprintf("%v", result["username"]),
		Email:          fmt.Sprintf("%v", result["email"]),
		Role:           fmt.Sprintf("%v", result["role"]),
		WhatsappNumber: fmt.Sprintf("%v", result["whatsappNumber"]),
		HashedPassword: fmt.Sprintf("%v", result["hashedPassword"]),
		ProfileImage:   fmt.Sprintf("%v", result["profileImage"]),
		FriendList:     friendList,
		LookingFor:     lookingFor,
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
				fmt.Println(err)
				next.ServeHTTP(w, r)
				return
			}

			//return user data to req
			ctx := context.WithValue(r.Context(), userCtxKey, user)
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
