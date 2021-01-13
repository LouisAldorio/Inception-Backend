package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFriendDetailByUsername(username string) []*model.Friend {
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	findOptions := options.FindOneOptions{}

	res := collection.FindOne(ctx, bson.M{
		"username": username,
	}, &findOptions)

	var comodityJSON bson.M
	if err := res.Decode(&comodityJSON); err != nil {
		fmt.Println(err)
	}

	//friend array
	var friends []*string
	for _, v := range comodityJSON["friendList"].(primitive.A) {
		friend := fmt.Sprintf("%v", v)
		friends = append(friends, &friend)
	}

	var result []*model.Friend
	for _, v := range friends {
		temp := &model.Friend{
			Username: *v,
		}
		result = append(result, temp)
	}

	return result
}
