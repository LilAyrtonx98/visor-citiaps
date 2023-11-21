package models

import (
	"context"
	"time"

	"github.com/citiaps/visor-mapas-backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GeoserverData allows to work with GeoServer's layers
type GeoserverData struct {
	Service			  string `json:"service,omitempty"`
	Version 		  string `json:"version,omitempty"`
	Request 	   	  string `json:"request,omitempty"`
	MaxFeatures 	  string `json:"max_features,omitempty"`
	OutputFormat 	  string `json:"output_format,omitempty"`
	Filename          string `json:"filename,omitempty"`
	CoordinatesSystem string `json:"coordinates_system,omitempty"`
	Workspace         string `json:"workspace,omitempty"`
	Datastore         string `json:"datastore,omitempty"`
}

//Url provides a strict separation of the service url used
type ParsedUrl struct {
	Protocol 	string		`json:"protocol,omitempty"`
	Host 		string		`json:"host,omitempty"`
	Port 		string		`json:"port,omitempty"` //If is empty an 80 is saved
	Path 		string		`json:"path,omitemtpy"`
}

//Provider contains data to interact with others systems layers
type Provider struct {
	Name 			string 			`json:"name,omitempty"` //Posible values "file, arcgis, geoserver" it's saved in lower case
	Url				string			`json:"url,omitempty"`
	ParsedUrl	 	ParsedUrl 		`json:"parsed_url,omitempty"`
	GeoserverData	GeoserverData	`json:"geoserverdata,omitempty"`
}

//Layer contains data about 
type Layer struct {
	ID            	primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	IDCategory    	primitive.ObjectID	`json:"id_category" bson:"id_category"`
	Name          	string             	`json:"name,omitempty"`
	Desc          	string				`json:"desc,omitempty"`
	Provider 		Provider  			`json:"provider,omitempty" bson:"provider,omitempty"`
	CreatedAt     	time.Time          	`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt		time.Time			`json:"updated_at" bson:"updated_at"`
}

func GetLayerByID(id primitive.ObjectID) Layer {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("layers")

	//Crea filtro de búsqueda para encontrar el documento por su id
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	//Variable donde se almacena el documento encontrado
	var result Layer

	//Buscar documento
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}

	return result
}