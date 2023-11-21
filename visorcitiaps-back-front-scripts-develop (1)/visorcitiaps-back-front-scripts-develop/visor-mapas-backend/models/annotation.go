package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//GeoJSON represent the coordinates of the annotation. Check: CRS 4326
type GeoJSON struct {
	GeoJSONType string    `json:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
}

//Annotation struct allows users annotate on a map, allows sharing these annotation with a group
type Annotation struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IDUser    primitive.ObjectID `json:"id_user,omitempty" bson:"_id_user,omitempty"`
	IDMap     primitive.ObjectID `json:"id_map,omitempty" bson:"_id_map,omitempty"`
	IDGroup   primitive.ObjectID `json:"id_group" bson:"_id_group"`
	Text      string             `json:"text,omitempty" bson:"text,omitempty"`
	IsShared  *bool 			 `json:"is_shared" bson:"_is_shared,omitempty"`
	Location  GeoJSON            `json:"location,omitempty" bson:"location,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time			 `json:"updated_at" bson:"updated_at"`
}