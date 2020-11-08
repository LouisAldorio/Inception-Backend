// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        *User  `json:"user"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUser struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Role            string `json:"role"`
	ConfirmPassword string `json:"confirm_password"`
}

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	HashedPassword string `json:"hashed_password"`
}

type UserOps struct {
	Register string         `json:"register"`
	Login    *LoginResponse `json:"login"`
}
