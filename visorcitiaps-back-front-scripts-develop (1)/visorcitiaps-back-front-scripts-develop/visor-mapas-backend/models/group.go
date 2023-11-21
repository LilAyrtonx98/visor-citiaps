package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//type JSON = map[string]interface{}

//Group corresponde a un grupo de usuarios
type Group struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Desc      string             `json:"desc,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	//UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

/*
// Serialize serializes post data
func (g Group) Serialize() JSON {
	return JSON{
		"id":         g.ID,
		"name":       g.Name,
		"desc":       g.Desc,
		"created_at": g.CreatedAt,
		"updated_at": g.UpdatedAt,
	}
}
*/
