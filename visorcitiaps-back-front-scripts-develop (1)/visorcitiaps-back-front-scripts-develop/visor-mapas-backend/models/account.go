package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NewPassword string             `json:"new_password,omitempty" bson:"new_password,omitempty"`
	OldPassword string             `json:"old_password,omitempty" bson:"old_password,omitempty"`
}
