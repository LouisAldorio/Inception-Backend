package service

import (
	"fmt"
	"myapp/graph/model"
	"myapp/utils"

	"context"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

//Register User
func Register(ctx context.Context, input model.NewUser) (string, error) {
	isValid, _ := utils.ValidateInput(ctx, input)
	if !isValid {
		return "", gqlerror.Errorf("%s", "input validation error")
	}

	user := &model.User{
		Email:          input.Email,
		Role:           input.Role,
		Username:       input.Username,
		HashedPassword: utils.HashPassword(input.Password),
	}

	// client := config.MongodbConnect()
	// collection := client.Database("Inception").Collection("Users")

	//insert to mongoDB
	fmt.Println(user)

	return fmt.Sprintf("%s", "Registration success"), nil
}

func Login(ctx context.Context, input model.LoginUser) (*model.LoginResponse, error) {
	var user model.User

	// client := config.MongodbConnect()
	// collection := client.Database("Inception").Collection("Users")

	//get user from mongoDB where username

	//dummy data
	// user = model.User{
	// 	Username: "felixyangsen",
	// 	Email: "felix@inception.com",
	// 	Role: "reseller",
	// 	HashedPassword: "$2a$04$Z5soW8sJwwch2NjbwnV4aeTR.aEnzHHOI9DtMq/RRITxX3d0JWbzS",
	// }

	if user.Username == "" {
		return nil, gqlerror.Errorf("%s", "user not found")
	}

	if isValid := utils.CheckPassword(user.HashedPassword, input.Password); !isValid {
		return nil, gqlerror.Errorf("%s", "Incorrect username or password")
	}

	accessToken, _ := utils.CreateToken(user)

	return &model.LoginResponse{
		AccessToken: accessToken,
		User:        &user,
	}, nil
}
