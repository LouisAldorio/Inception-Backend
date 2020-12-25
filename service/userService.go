package service

import (
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
func Register(ctx context.Context, input model.NewUser) (*model.LoginResponse, error) {
	if isValid, err := utils.ValidateInput(ctx, input); !isValid {
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
		{"hashedPassword", utils.HashPassword(input.Password)},
		{"role", input.Role},
		{"whatsappNumber", input.WhatsappNumber},
		{"profileImage","https://www.baytekent.com/wp-content/uploads/2016/12/facebook-default-no-profile-pic1.jpg"},
		{"friendList",[]string{}},
	})
	if err != nil {
		return nil, gqlerror.Errorf("Registration failed %s", err.Error())
	}

	user := model.User{
		Username:       input.Username,
		Email:          input.Email,
		Role:           input.Role,
		WhatsappNumber: *input.WhatsappNumber,
	}
	accessToken, _ := utils.CreateToken(user)

	return &model.LoginResponse{
		AccessToken: accessToken,
		User:        &user,
	}, nil
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

	if isValid := utils.CheckPassword(userJSON["hashedPassword"].(string), input.Password); !isValid {
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
