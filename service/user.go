package service

import (
	"fmt"
	"log"
	"myapp/graph/model"
	"myapp/utils"

	"context"
)

//Register User
func Register(ctx context.Context, input model.NewUser) *model.User {

	temp := utils.ForContext(ctx)
	fmt.Println(temp)

	var user *model.User
	// client := config.MongodbConnect()
	// collection := client.Database("Inception").Collection("Users")

	isValid, err := utils.ValidateInput(ctx, input)
	if err != nil {
		fmt.Println(err)
	}

	if isValid {
		token, err := utils.CreateToken(input.Username, input.Password)
		if err != nil {
			log.Println(err)
		}
		user = &model.User{
			Email:    input.Email,
			Password: input.Password,
			Role:     input.Role,
			Username: input.Username,
			Token:    token,
		}
	}
	return user
}


