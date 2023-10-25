package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	ID   primitive.ObjectID `json:"_id,omitempty"`
	Name string             `json:"name,omitempty" validate:"required"`
	Age  int                `json:"age,omitempty" validate:"required"`
}
