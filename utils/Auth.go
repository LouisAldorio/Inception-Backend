package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CreateToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", gqlerror.Errorf(fmt.Sprintf("%s", err))
	}
	return tokenString, nil
}

func ValidateToken(t string) (string, error) {
	if t == "" {
		return "", gqlerror.Errorf("Authorization token must be present")
	}

	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := Claims["username"].(string)
		return username, nil
	} else {
		return "", gqlerror.Errorf("Invalid authorization token")
	}
}
