package utils

import (
	"context"
	"myapp/graph/model"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

func ValidateInput(ctx context.Context, input model.NewUser)(bool,error) {

	totalErrors := 0

	if len(strings.TrimSpace(input.Email)) == 0 {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Email must not be empty!",
		})
		totalErrors++
	}

	if len(strings.TrimSpace(input.Username)) == 0 {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Username must not be empty!",
		})
		totalErrors++
	}

	if len(strings.TrimSpace(input.Password)) == 0 {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Password must not be empty!",
		})
		totalErrors++
	}

	if input.Password != input.ConfirmPassword {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Confirm password doesn't match!",
		})
		totalErrors++
	}

	if totalErrors > 0{
		return false,nil
	}

	return true,nil
}
