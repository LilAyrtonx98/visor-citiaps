package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Categoría corresponde a un grupo de usuarios
type Map struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Desc      string             `json:"desc,omitempty"`
	ImgURL    string             `json:"imgurl,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	//UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
