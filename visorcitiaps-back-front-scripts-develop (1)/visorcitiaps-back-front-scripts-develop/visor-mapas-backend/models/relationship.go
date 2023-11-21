package models

import (
	"context"
	"log"
	"time"

	"github.com/citiaps/visor-mapas-backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//RelationshipMapUseres la relación entre mapa y usuario (N:M)
type RelationshipMapUser struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IDMap     primitive.ObjectID `json:"id_user" bson:"id_map"`
	IDUser    primitive.ObjectID `json:"id_layer" bson:"id_user"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

//RelationshipUserLayer es la relación entre mapa y geoproceso (N:M)
type RelationshipMapGeoprocessing struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IDMap           primitive.ObjectID `json:"id_user" bson:"id_map"`
	IDGeoprocessing primitive.ObjectID `json:"id_geoprocessing" bson:"id_geoprocessing"`
	CreatedAt       time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

//RelationshipMapLayer es la relación entre mapa y capa (N:M)
type RelationshipMapLayer struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IDMap     primitive.ObjectID `json:"id_user" bson:"id_map"`
	IDLayer   primitive.ObjectID `json:"id_layer" bson:"id_layer"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

//Borra todas las relaciones maps_layers al borrar una capa
func DeleteAllRelMapsByLayer(idLayer primitive.ObjectID) bool {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("maps_layers")

	filter := bson.D{primitive.E{Key: "id_layer", Value: idLayer}}

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//Borra todas las relaciones maps_geoprocessings al borrar un geoproceso
func DeleteAllRelMapsByGeoprocessings(idGeoprocessing primitive.ObjectID) bool {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("maps_geoprocessings")

	filter := bson.D{primitive.E{Key: "id_geoprocessing", Value: idGeoprocessing}}

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//Borra todas las relaciones maps_users al borrar un usuario
func DeleteAllRelMapsByUsers(idUser primitive.ObjectID) bool {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("maps_users")

	filter := bson.D{primitive.E{Key: "id_user", Value: idUser}}

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//Borra todas las relaciones users_geoprocessings y users_layers al borrar un usuario
func DeleteAllRelByMap(idMap primitive.ObjectID) bool {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("maps_users")

	filter := bson.D{primitive.E{Key: "id_map", Value: idMap}}

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return false
	}

	collection2 := db.Connection.Collection("maps_geoprocessings")

	_, err = collection2.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return false
	}

	collection3 := db.Connection.Collection("maps_layers")

	_, err = collection3.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
