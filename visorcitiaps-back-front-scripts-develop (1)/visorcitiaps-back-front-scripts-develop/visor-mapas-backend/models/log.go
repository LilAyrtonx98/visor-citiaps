package models

import (
	"context"
	"log"
	"time"

	"github.com/citiaps/visor-mapas-backend/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IDUserExecutor     string             `json:"id_user_executor" bson:"id_user_executor,omitempty"`
	Action             string             `json:"action,omitempty"`
	Collection         string             `json:"collection,omitempty"`
	IDResource         string             `json:"id_resource" bson:"id_resource,omitempty"`
	NewResource        string             `json:"new_resource" bson:"new_resource,omitempty"`
	IDUserAffected     string             `json:"id_user_affected" bson:"id_user_affected,omitempty"`
	NotifyUserAffected bool               `json:"notify_user_affected,omitempty" bson:"notify_user_affected,omitempty"`
	CreatedAt          time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func NewLog(userEx, action, collection, resource, newResource, userAf string, notify bool) *Log {
	l := new(Log)
	l.IDUserExecutor = userEx
	l.IDResource = resource
	l.IDUserAffected = userAf
	l.NewResource = newResource
	l.Action = action
	l.Collection = collection
	l.NotifyUserAffected = notify
	//Establecer fecha de creación del documento
	l.CreatedAt = time.Now()
	return l
}

//InsertOneUser agrega un documento usuario
func (l *Log) SaveLog() bool {

	//Acceder a la colección donde están los usuarios
	collection := db.Connection.Collection("logs")

	//Insertar documento
	_, err := collection.InsertOne(context.TODO(), l)
	if err != nil {
		log.Println("No se pudo registrar acción de usuario")
		return false
	}

	return true
}

func CreateLog(userEx, action, collection, resource, newResource, userAf string, notify bool) bool {
	// Si no se encuentra el usuario que ejecuta la acción, no guardar log
	if userEx == "" {
		return false
	}
	var value bool
	go func() {
		value = NewLog(userEx, action, collection, resource, newResource, userAf, notify).SaveLog()
	}()
	return value
}
