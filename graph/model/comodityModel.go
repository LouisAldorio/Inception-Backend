package model

import "go.mongodb.org/mongo-driver/mongo"

type Comodity struct {
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	UnitPrice   float64 `json:"unit_price"`
	UnitType    string  `json:"unit_type"`
	MinPurchase string  `json:"min_purchase"`
	Description *string `json:"description"`
	User        *User   `json:"user"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func (m *Comodity) decodeMongo(c *mongo.Cursor) {

}
