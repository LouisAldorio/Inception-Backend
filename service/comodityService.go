package service

import (
	"context"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"time"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ComodityGetList(ctx context.Context, limit, page *int) ([]*model.Comodity, error) {
	var comodities []*model.Comodity

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	//get user from mongoDB where username
	findOptions := options.FindOptions{}

	res, err := collection.Find(ctx, bson.M{}, &findOptions)
	if err != nil {
		return nil, err
	}

	for res.Next(ctx) {
		var comodityJSON bson.M
		if err := res.Decode(&comodityJSON); err != nil {
			return nil, err
		}

		var temp = comodityJSON["description"].(string)
		var user model.User

		var userJSON = comodityJSON["user"]
		mapstructure.Decode(userJSON, &user)

		comodity := model.Comodity{
			Name:        comodityJSON["name"].(string),
			UnitType:    comodityJSON["unit_type"].(string),
			Description: &temp,
			MinPurchase: comodityJSON["min_purchase"].(string),
			User:        &user,
		}

		comodities = append(comodities, &comodity)
	}

	return comodities, nil
}

func CommodityCreate(input *model.NewComodity,user *model.User) *model.Comodity{
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Comodities")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.D{
		{"name",input.Name},
		{"minPurchase",input.MinPurchase},
		{"unitPrice",input.UnitPrice},
		{"unitType",input.UnitType},
		{"description",input.Description},
		{"images",input.Images},
		{"user",user},
	})

	if err != nil {
		log.Println(err)
	}

	result := model.Comodity{
		Name: input.Name,
		Image: input.Images,
		UnitPrice: input.UnitPrice,
		UnitType: input.UnitType,
		MinPurchase: input.MinPurchase,
		Description: &input.Description,
		User: user,
	}
	return &result
}
