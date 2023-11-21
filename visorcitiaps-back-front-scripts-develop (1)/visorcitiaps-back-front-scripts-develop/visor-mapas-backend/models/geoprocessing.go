package models

import (
	"context"
	"time"

	"github.com/citiaps/visor-mapas-backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Layer es una capa almacenada en la aplicación
type Geoprocessing struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Desc      string             `json:"desc,omitempty"`
	GeoURL    string             `json:"geo_url,omitempty" bson:"geo_url,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func GetGeoprocessingByID(id primitive.ObjectID) Geoprocessing {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("geoprocessings")

	//Crea filtro de búsqueda para encontrar el documento por su id
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	//Variable donde se almacena el documento encontrado
	var result Geoprocessing

	//Buscar documento
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}

	return result
}
