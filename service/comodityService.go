package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"

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
			Image:       comodityJSON["image"].(string),
			UnitPrice:   comodityJSON["unit_price"].(float64),
			UnitType:    comodityJSON["unit_type"].(string),
			Description: &temp,
			MinPurchase: comodityJSON["min_purchase"].(string),
			User:        &user,
			// CreatedAt:   comodityJSON["created_at"].(string),
			// UpdatedAt:   comodityJSON["updated_at"].(*string),
		}

		comodities = append(comodities, &comodity)
	}

	return comodities, nil
}
