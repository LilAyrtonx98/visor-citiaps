package models

import (
	"context"
	"time"

	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Permissions represented by boolean that a User has.
type Permissions struct {
	Users      bool `json:"users,omitempty"`
	Layers     bool `json:"layers,omitempty" bson:"layers"`
	Geo        bool `json:"geo,omitempty" bson:"geo"`
	Maps       bool `json:"maps,omitempty"`
	Visor      bool `json:"visor,omitempty"`
	Annotation bool `json:"annotation,omitempty"`
}

//User representa a un usuario de la aplicación
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IDGroup     primitive.ObjectID `json:"id_group" bson:"id_group"`
	Firstname   string             `json:"firstname,omitempty"`
	Lastname    string             `json:"lastname,omitempty"`
	Username    string             `json:"username,omitempty"`
	Password    string             `json:"password,omitempty"`
	Permissions Permissions        `json:"permissions,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func CheckPasswordByUsername(username string, password string) (bool, *User) {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("users")

	//Crea filtro de búsqueda para encontrar el documento por su username
	filter := bson.M{"username": username}

	//Variable donde se almacena el documento encontrado
	var result User

	//Buscar documento
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false, nil
	}

	//Verificar password credencial con hash en db
	isAuth := utils.CheckPasswordHash(password, result.Password)

	return isAuth, &result
}

func CheckPasswordByID(id primitive.ObjectID, password string) (bool, *User) {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("users")

	//Crea filtro de búsqueda para encontrar el documento por su ID
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	//Variable donde se almacena el documento encontrado
	var result User

	//Buscar documento
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false, nil
	}

	//Verificar password credencial con hash en db
	isAuth := utils.CheckPasswordHash(password, result.Password)

	return isAuth, &result
}

func GetUsernameByID(id primitive.ObjectID) string {
	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("users")

	//Crea filtro de búsqueda para encontrar el documento por su id
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	//Variable donde se almacena el documento encontrado
	var result User

	//Buscar documento
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return ""
	}

	return result.Username
}

func NotifyUserAssigedResource(idUser, idResource primitive.ObjectID, resourceType string) {
	go func() {
		toMail := GetUsernameByID(idUser)
		switch resourceType {
		case "layer":
			resource := GetLayerByID(idResource)
			messageBody := "La capa " + resource.Name + " se te ha asignado en el Visor"
			utils.CreateMail([]string{toMail}, "Asignación de capa", messageBody)
		case "geoprocessing":
			resource := GetGeoprocessingByID(idResource)
			messageBody := "El geoproceso " + resource.Name + " se te ha asignado en el Visor"
			utils.CreateMail([]string{toMail}, "Asignación de geoproceso", messageBody)
		default:
			return
		}
	}()
}
