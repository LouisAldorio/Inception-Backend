package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"myapp/utils"
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

func UpdateAddedOrRemoveUserFriendlist(username string,friend string, remove bool){
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userDetail,_ := GetUserByUsername(username) 
	filter := bson.M{"username": username}

	newFriendArray := []string{}

	var update primitive.M
	if !remove {
		for _,v := range userDetail.FriendList{
			newFriendArray = append(newFriendArray, v)
		}
		newFriendArray = append(newFriendArray, friend)
		update = bson.M{
			"$set": bson.M{"friendList": newFriendArray},
		}
	}else {
		fmt.Println(userDetail.FriendList,friend)
		for _,v := range userDetail.FriendList{
			if v != friend{
				newFriendArray = append(newFriendArray, v)
			}
		}
		fmt.Println(newFriendArray)
		update = bson.M{
			"$set": bson.M{"friendList": newFriendArray},
		}
	}
	
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		fmt.Println(result.Err())
	}
}

func AddFriend(friends []*string,username string) *model.User{
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//update to friend
	userDetail,_ := GetUserByUsername(username) 
	prevFriendList := userDetail.FriendList

	nonPointerFriends := []string{}
	for _,v := range friends{
		nonPointerFriends = append(nonPointerFriends, *v)
	}


	if len(prevFriendList) < len(nonPointerFriends){
		//add friend
		newFriend := utils.Difference(nonPointerFriends,prevFriendList)
		UpdateAddedOrRemoveUserFriendlist(newFriend[0],username,false)

	}else if len(prevFriendList) > len(nonPointerFriends){
		//remove friend
		removedFriend := utils.Difference(prevFriendList,nonPointerFriends)
		UpdateAddedOrRemoveUserFriendlist(removedFriend[0],username,true)
	}

	//update my own friendList
	filter := bson.M{"username": username}

	update := bson.M{
		"$set": bson.M{"friendList": friends},
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		fmt.Println(result.Err())
	}

	// 9) Decode the result
	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	if decodeErr != nil {
		fmt.Println(decodeErr)
	}

	var friendList []string
	for _, v := range doc["friendList"].(primitive.A) {
		friend := fmt.Sprintf("%v", v)
		friendList = append(friendList, friend)
	}

	var lookingFor []string
	for _, v := range doc["lookingFor"].(primitive.A) {
		item := fmt.Sprintf("%v", v)
		lookingFor = append(lookingFor, item)
	}

	return &model.User{
		Username:       fmt.Sprintf("%v", doc["username"]),
		Email:          fmt.Sprintf("%v", doc["email"]),
		Role:           fmt.Sprintf("%v", doc["role"]),
		WhatsappNumber: fmt.Sprintf("%v", doc["whatsappNumber"]),
		HashedPassword: fmt.Sprintf("%v", doc["hashedPassword"]),
		ProfileImage:   fmt.Sprintf("%v", doc["profileImage"]),
		FriendList:     friendList,
		LookingFor:     lookingFor,
	}
}
