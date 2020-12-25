package service

import (
	"context"
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"time"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ComodityGetList(ctx context.Context, limit, page *int) ([]*model.Comodity, error) {
	var comodities []*model.Comodity

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	//get user from mongoDB where username
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})

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
		var user model.User

		var userJSON = comodityJSON["user"]
		mapstructure.Decode(userJSON, &user)

		comodity := model.Comodity{
			ID:          comodityJSON["_id"].(primitive.ObjectID).String(),
			Name:        comodityJSON["name"].(string),
			UnitType:    comodityJSON["unitType"].(string),
			Description: &temp,
			Image:       images,
			UnitPrice:   comodityJSON["unitPrice"].(string),
			MinPurchase: comodityJSON["minPurchase"].(string),
			User:        &user,
		}

		comodities = append(comodities, &comodity)
	}

	return comodities, nil
}

func CommodityCreate(input *model.NewComodity, user *model.User) *model.Comodity {
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.D{
		{"name", input.Name},
		{"minPurchase", input.MinPurchase},
		{"unitPrice", input.UnitPrice},
		{"unitType", input.UnitType},
		{"description", input.Description},
		{"images", input.Images},
		{"user", user},
	})

	if err != nil {
		log.Println(err)
	}

	result := model.Comodity{
		Name:        input.Name,
		Image:       input.Images,
		UnitPrice:   input.UnitPrice,
		UnitType:    input.UnitType,
		MinPurchase: input.MinPurchase,
		Description: &input.Description,
		User:        user,
	}
	return &result
}
