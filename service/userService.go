package service

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"myapp/utils"
	"time"

	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	var err error
	if input.Role == "Distributor" {
		_, err = collection.InsertOne(ctx, bson.D{
			{"email", input.Email},
			{"username", input.Username},
			{"hashedPassword", utils.HashPassword(input.Password)},
			{"role", input.Role},
			{"whatsappNumber", input.WhatsappNumber},
			{"profileImage", "https://www.baytekent.com/wp-content/uploads/2016/12/facebook-default-no-profile-pic1.jpg"},
			{"friendList", []string{}},
			{"lookingFor", []string{}},
		})
	} else if input.Role == "Supplier" {
		_, err = collection.InsertOne(ctx, bson.D{
			{"email", input.Email},
			{"username", input.Username},
			{"hashedPassword", utils.HashPassword(input.Password)},
			{"role", input.Role},
			{"whatsappNumber", input.WhatsappNumber},
			{"profileImage", "https://www.baytekent.com/wp-content/uploads/2016/12/facebook-default-no-profile-pic1.jpg"},
			{"friendList", []string{}},
			{"lookingFor", []string{}},
		})
	}

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

func GetUserByRole(role string) []*model.User {
	var users []*model.User

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOptions := options.FindOptions{}
	cur, err := collection.Find(ctx, bson.D{
		{"role", role},
	}, &findOptions)

	for cur.Next(ctx) {
		var result bson.M
		err = cur.Decode(&result)
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

		user := &model.User{
			Username:       fmt.Sprintf("%v", result["username"]),
			Email:          fmt.Sprintf("%v", result["email"]),
			Role:           fmt.Sprintf("%v", result["role"]),
			WhatsappNumber: fmt.Sprintf("%v", result["whatsappNumber"]),
			HashedPassword: fmt.Sprintf("%v", result["hashedPassword"]),
			ProfileImage:   fmt.Sprintf("%v", result["profileImage"]),
			FriendList:     friendList,
			LookingFor:     lookingFor,
		}
		users = append(users, user)
	}

	return users
}

func UpdateUserProfile(user *model.User, input model.EditUser)*model.User{
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{"username", user.Username}}
    update := bson.D{{"$set",
        bson.D{
			{"email", input.Email},
			{"whatsappNumber",input.WhatsappNumber},
			{"profileImage",input.ProfileImage},
			{"lookingFor",input.LookingFor},
        },
	}}
	
	_, err := collection.UpdateOne(
        ctx,
        filter,
        update,
    )
    if err != nil {
        fmt.Println(err)
	}

	result,_ := GetUserByUsername(user.Username)
	
	return result
}

func GetUserByWhereIn(usernames []*string)[]*model.User {
	var result []*model.User

	for _,v := range usernames{
		temp,_ := GetUserByUsername(*v)
		result = append(result, temp)
	}

	return result
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
