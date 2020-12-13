package service

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"myapp/utils"

	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Register User
func Register(ctx context.Context, input model.NewUser) (*string, error) {
	if isValid, err := utils.ValidateInput(ctx, input); isValid {
		fmt.Println(err)
		return nil, err
	}

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	findOptions := options.FindOneOptions{}

	cur := collection.FindOne(ctx, bson.M{"username": input.Username}, &findOptions)
	if cur.Err() == nil {
		return nil, gqlerror.Errorf("%s", "User has been registered!")
	}

	_, err := collection.InsertOne(ctx, bson.D{
		{"email", input.Email},
		{"username", input.Username},
		{"password", utils.HashPassword(input.Password)},
		{"role", input.Role},
		{"whatsapp_number", input.WhatsappNumber},
	})
	if err != nil {
		fmt.Println(err)
		return nil, gqlerror.Errorf("Registration failed %s", err.Error())
	}

	response := fmt.Sprintf("%s", "Registration success")

	return &response, nil
}

func Login(ctx context.Context, input model.LoginUser) (*model.LoginResponse, error) {
	var userJSON bson.M

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	//get user from mongoDB where username
	findOptions := options.FindOneOptions{}

	cur := collection.FindOne(ctx, bson.M{"username": input.Username}, &findOptions)
	if cur.Err() != nil {
		return nil, gqlerror.Errorf("%s", "User not found!")
	}
	err := cur.Decode(&userJSON)
	if err != nil {
		log.Fatal(err)
	}

	if isValid := utils.CheckPassword(userJSON["password"].(string), input.Password); !isValid {
		return nil, gqlerror.Errorf("%s", "Incorrect username or password")
	}

	var user model.User
	mapstructure.Decode(userJSON, &user)

	accessToken, _ := utils.CreateToken(user)

	return &model.LoginResponse{
		AccessToken: accessToken,
		User:        &user,
	}, nil
}
