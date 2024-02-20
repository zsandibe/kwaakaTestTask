package domain

import "time"

type Weather struct {
	Id          string    `json:"id" bson:"id"`
	City        string    `json:"city" bson:"city"`
	Temperature float64   `json:"temperature" bson:"temperature"`
	CreatedAt   time.Time `json:"created_ad" bson:"created_ad"`
	UpdatedAt   time.Time `json:"updated_ad" bson:"updated_ad"`
}
