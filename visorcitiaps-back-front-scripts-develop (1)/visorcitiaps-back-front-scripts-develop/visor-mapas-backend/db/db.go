package db

import (
	"context"
	"fmt"
	"log"

	"github.com/citiaps/visor-mapas-backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connection pool a la base de datos
var Connection *mongo.Database

//Setup realiza la conexión con DB. Solo se llama una vez en el main
func Setup() {
	//Conectar a MongoDB
	uri := fmt.Sprintf(`mongodb://%s:%s@%s:%s/%s`,
		utils.Config.Database.User,
		utils.Config.Database.Pass,
		utils.Config.Database.Host,
		utils.Config.Database.Port,
		utils.Config.Database.Name)
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	//Verificar conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Println("Conectado a MongoDB")
	Connection = client.Database(utils.Config.Database.Name)
}
