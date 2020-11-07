package service

import (
	"fmt"
	"myapp/config"
)

func Register() *string{
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Users")

	fmt.Println(collection)
	return nil
}