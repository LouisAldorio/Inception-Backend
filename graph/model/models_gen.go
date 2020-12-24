// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CommodityOps struct {
	Create *Comodity `json:"create"`
}

type Comodity struct {
	Name        string    `json:"name"`
	Image       []*string `json:"image"`
	UnitPrice   string    `json:"unit_price"`
	UnitType    string    `json:"unit_type"`
	MinPurchase string    `json:"min_purchase"`
	Description *string   `json:"description"`
	User        *User     `json:"user"`
}

type ComodityPagination struct {
	Limit     *int        `json:"limit"`
	Page      *int        `json:"page"`
	TotalItem int         `json:"total_item"`
	Nodes     []*Comodity `json:"nodes"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        *User  `json:"user"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewComodity struct {
	Name        string    `json:"name"`
	MinPurchase string    `json:"min_purchase"`
	UnitType    string    `json:"unit_type"`
	UnitPrice   string    `json:"unit_price"`
	Description string    `json:"description"`
	Images      []*string `json:"images"`
}

type NewUser struct {
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	Role            string  `json:"role"`
	WhatsappNumber  *string `json:"whatsapp_number"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
}

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	WhatsappNumber string `json:"whatsapp_number"`
	HashedPassword string `json:"hashed_password"`
}

type UserOps struct {
	Register *LoginResponse `json:"register"`
	Login    *LoginResponse `json:"login"`
}
