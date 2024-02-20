package domain

import "time"

type Weather struct {
	Id          string    `json:"id" bson:"id"`
	City        string    `json:"city" bson:"city"`
	Temperature float32   `json:"temperature" bson:"temperature"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
