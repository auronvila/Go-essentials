package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Movie   string             `json:"movie"`
	Watched bool               `json:"watched"`
}
