package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Account = models.Account

func FindMyUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("users")

		idString, ok := auth.GetUserIDFromCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Obtener id del documento a buscar en el id string
		id, err := primitive.ObjectIDFromHex(idString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Variable donde se almacena el documento encontrado
		var result User

		//Buscar documento
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		result.Password = ""

		//Retornar documento
		c.JSON(http.StatusOK, result)
	}
}

func ChangeMyPassword() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("users")

		//De request body a User
		var elem Account
		c.BindJSON(&elem)

		//Establecer fecha de actualización del documento
		//elem.UpdatedAt = time.Now()

		idString, ok := auth.GetUserIDFromCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Obtener id del documento a actualizar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(idString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Verificar contraeña anterior coincide con usuario
		if ok, _ := models.CheckPasswordByID(id, elem.OldPassword); !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Hash password
		password := elem.NewPassword
		hassPassword, err := utils.HashPassword(password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Actualizar contraseña
		elem.NewPassword = hassPassword

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Valores a actualizar del documento
		update := bson.M{
			"$set": bson.M{
				"password": elem.NewPassword,
				//"updated_at": elem.UpdatedAt,
			},
		}

		//Actualizar documento
		_, err = collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		models.CreateLog(idEx, "Cambia contraseña", "", "", "", "", false)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento actualizado exitosamente",
		})
	}
}

func CheckUniqueUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("users")

		//De request body a User
		var elem User
		c.BindJSON(&elem)

		//Crea filtro de búsqueda para encontrar el documento por su username
		filter := bson.M{"username": elem.Username}

		//Variable donde se almacena el documento encontrado
		var result *User

		//Buscar documento
		_ = collection.FindOne(context.TODO(), filter).Decode(&result)

		//Si hay un resultado, el username ya está registrado
		if result != nil {
			c.JSON(http.StatusOK, gin.H{
				"valid":   false,
				"message": elem.Username + " ya se encuentra registrado",
			})
			return
		}

		//Si no hay resultados, username disponible
		c.JSON(http.StatusOK, gin.H{
			"valid":   true,
			"message": elem.Username + " está disponible",
		})
	}
}

func FindMyLogs() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("logs")

		//Paginación
		findOptions, err := utils.PaginationFindOptions(c.DefaultQuery("page", utils.Config.APIRest.Page), c.DefaultQuery("size", utils.Config.APIRest.Size))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		idString, ok := auth.GetUserIDFromCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Todas las capas con id de categoria igual al consultado
		filter := bson.M{"id_user_executor": idString}

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Log

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Iterar sobre el cursor para guardar los documentos en el arreglo results
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem Log
			err := cur.Decode(&elem)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible actualizar el documento"})
				return
			}
			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Cerrar el cursor
		cur.Close(context.TODO())

		//Obtener cantidad de páginas de la colección (para paginación)
		//Se salta desde página 1 hasta page-1
		count, _ := collection.CountDocuments(context.TODO(), filter, options.Count().SetSkip(*findOptions.Skip))
		totalPages := utils.EstimatedTotalPages(c.DefaultQuery("size", utils.Config.APIRest.Size), count, *findOptions.Skip)
		c.Writer.Header().Set("TotalPages", totalPages)

		//Retornar documentos
		c.JSON(http.StatusOK, results)

	}
}

func FindMyNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("logs")

		//Paginación
		findOptions, err := utils.PaginationFindOptions(c.DefaultQuery("page", utils.Config.APIRest.Page), c.DefaultQuery("size", utils.Config.APIRest.Size))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		idString, ok := auth.GetUserIDFromCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Todas las capas con id de categoria igual al consultado
		filter := bson.M{
			"id_user_affected":     idString,
			"notify_user_affected": true,
		}

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Log

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Iterar sobre el cursor para guardar los documentos en el arreglo results
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem Log
			err := cur.Decode(&elem)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible actualizar el documento"})
				return
			}
			log.Println(elem.Collection)
			switch elem.Collection {
			case "layers":
				log.Println("CAPAS")
				hexID, _ := primitive.ObjectIDFromHex(elem.IDResource)
				aux := models.GetLayerByID(hexID)
				elem.IDResource = aux.Name
				break
			case "users_layers":
				log.Println("USUARIOS_CAPAS")
				hexID, _ := primitive.ObjectIDFromHex(elem.IDResource)
				aux := models.GetLayerByID(hexID)
				elem.IDResource = aux.Name
				break
			case "geoprocessings":
				log.Println("GEOPROCESOS")
				hexID, _ := primitive.ObjectIDFromHex(elem.IDResource)
				aux := models.GetGeoprocessingByID(hexID)
				elem.IDResource = aux.Name
				break
			case "users_geoprocessings":
				log.Println("USUARIOS_GEOPROCESOS")
				hexID, _ := primitive.ObjectIDFromHex(elem.IDResource)
				aux := models.GetGeoprocessingByID(hexID)
				elem.IDResource = aux.Name
				break
			default:
				break
			}

			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Cerrar el cursor
		cur.Close(context.TODO())

		//Obtener cantidad de páginas de la colección (para paginación)
		//Se salta desde página 1 hasta page-1
		count, _ := collection.CountDocuments(context.TODO(), filter, options.Count().SetSkip(*findOptions.Skip))
		totalPages := utils.EstimatedTotalPages(c.DefaultQuery("size", utils.Config.APIRest.Size), count, *findOptions.Skip)
		c.Writer.Header().Set("TotalPages", totalPages)

		//Retornar documentos
		c.JSON(http.StatusOK, results)

	}
}
