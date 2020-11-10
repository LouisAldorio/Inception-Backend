package service

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"myapp/utils"

	"context"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Register User
func Register(ctx context.Context, input model.NewUser) (string, error) {

	isValid, err := utils.ValidateInput(ctx, input)
	if !isValid {
		return "", err
	}

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	findOptions := options.FindOneOptions{}

	cur := collection.FindOne(ctx, bson.M{"username": input.Username}, &findOptions)
	if cur.Err() == nil {
		return "", gqlerror.Errorf("%s", "User has been registered!")
	}

	_, mongoInsertErr := collection.InsertOne(ctx, bson.D{
		{"email", input.Email},
		{"username", input.Username},
		{"password", utils.HashPassword(input.Password)},
		{"role", input.Role},
	})
	if mongoInsertErr != nil {
		log.Println(err)
	}

	return fmt.Sprintf("%s", "Registration success"), nil
}

func Login(ctx context.Context, input model.LoginUser) (*model.LoginResponse, error) {
	var user bson.M

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	//get user from mongoDB where username
	findOptions := options.FindOneOptions{}

	cur := collection.FindOne(ctx, bson.M{"username": input.Username}, &findOptions)
	if cur.Err() != nil {
		return nil, gqlerror.Errorf("%s", "User not found!")
	}
	err := cur.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	if isValid := utils.CheckPassword(user["password"].(string), input.Password); !isValid {
		return nil, gqlerror.Errorf("%s", "Incorrect username or password")
	}

	loggedInUser := model.User{
		Email:          user["email"].(string),
		HashedPassword: user["password"].(string),
		Role:           user["role"].(string),
		Username:       user["username"].(string),
	}

	accessToken, _ := utils.CreateToken(loggedInUser)

	return &model.LoginResponse{
		AccessToken: accessToken,
		User:        &loggedInUser,
	}, nil
}
