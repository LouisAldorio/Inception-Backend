package service

import (
	"context"
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTotalCommodity(ctx context.Context) (int, error) {
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func ComodityGetList(ctx context.Context, limit, page *int) ([]*model.Comodity, error) {
	var comodities []*model.Comodity

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	one := 1
	var offset int
	if limit != nil && page != nil {
		offset = *limit * (*page - one)
	}

	//get user from mongoDB where username
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(int64(*limit)).SetSkip(int64(offset))

	res, err := collection.Find(ctx, bson.M{}, &findOptions)
	if err != nil {
		return nil, err
	}

	for res.Next(ctx) {
		var comodityJSON bson.M
		if err := res.Decode(&comodityJSON); err != nil {
			return nil, err
		}

		//image array
		var images []*string
		for _, v := range comodityJSON["images"].(primitive.A) {
			image := fmt.Sprintf("%v", v)
			images = append(images, &image)
		}

		var temp = comodityJSON["description"].(string)

		comodity := model.Comodity{
			ID:          comodityJSON["_id"].(primitive.ObjectID).String(),
			Name:        comodityJSON["name"].(string),
			UnitType:    comodityJSON["unitType"].(string),
			Description: &temp,
			Image:       images,
			UnitPrice:   comodityJSON["unitPrice"].(string),
			MinPurchase: comodityJSON["minPurchase"].(string),
			Username:    comodityJSON["username"].(string),
		}

		comodities = append(comodities, &comodity)
	}

	return comodities, nil
}

func GetCommoditiesByUsername(ctx context.Context, username string) []*model.Comodity {
	var comodities []*model.Comodity

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	//get user from mongoDB where username
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})

	res, err := collection.Find(ctx, bson.M{
		"username": bson.M{"$eq": username},
	}, &findOptions)
	if err != nil {
		fmt.Println(err)
	}

	for res.Next(ctx) {
		var comodityJSON bson.M
		if err := res.Decode(&comodityJSON); err != nil {
			fmt.Println(err)
		}

		//image array
		var images []*string
		for _, v := range comodityJSON["images"].(primitive.A) {
			image := fmt.Sprintf("%v", v)
			images = append(images, &image)
		}

		var temp = comodityJSON["description"].(string)
		var user model.User

		var userJSON = comodityJSON["user"]
		mapstructure.Decode(userJSON, &user)

		comodity := model.Comodity{
			ID:          comodityJSON["_id"].(primitive.ObjectID).Hex(),
			Name:        comodityJSON["name"].(string),
			UnitType:    comodityJSON["unitType"].(string),
			Description: &temp,
			Image:       images,
			UnitPrice:   comodityJSON["unitPrice"].(string),
			MinPurchase: comodityJSON["minPurchase"].(string),
			Username:    comodityJSON["username"].(string),
		}

		comodities = append(comodities, &comodity)
	}

	return comodities
}

func CommodityCreate(ctx context.Context, input *model.NewComodity, username string) *model.Comodity {
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	cur, err := collection.InsertOne(ctx, bson.D{
		{"name", input.Name},
		{"minPurchase", input.MinPurchase},
		{"unitPrice", input.UnitPrice},
		{"unitType", input.UnitType},
		{"description", input.Description},
		{"images", input.Images},
		{"username", username},
	})

	if err != nil {
		log.Println(err)
	}

	result := model.Comodity{
		ID:          cur.InsertedID.(primitive.ObjectID).Hex(),
		Name:        input.Name,
		Image:       input.Images,
		UnitPrice:   input.UnitPrice,
		UnitType:    input.UnitType,
		MinPurchase: input.MinPurchase,
		Description: &input.Description,
		Username:    username,
	}
	return &result
}
