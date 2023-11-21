package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User = models.User

type CategoryWithLayers = models.CategoryWithLayers

//InsertOneUser agrega un documento usuario
func InsertOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("users")

		//De request body a User
		var elem User
		c.BindJSON(&elem)

		//Hash password
		password := elem.Password
		hassPassword, err := utils.HashPassword(password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//Actualizar usuario
		elem.Password = hassPassword

		//Establecer fecha de creación del documento
		elem.CreatedAt = time.Now()
		//Establecer fecha de actualización del documento
		//elem.UpdatedAt = time.Now()

		//Insertar documento
		insertResult, err := collection.InsertOne(context.TODO(), elem)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//Actualizar ID para response body
		elem.ID = insertResult.InsertedID.(primitive.ObjectID)

		elem.Password = ""

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Creación de usuario", "users", elem.ID.Hex(), string(serialized), "", false)

		//Retornar documento
		c.JSON(http.StatusCreated, elem)
	}
}

func FindOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("users")

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idUser"))
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

func FindAllUsers(pagination bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("users")

		findOptions := new(options.FindOptions)
		var err error

		if pagination {
			//Paginación
			findOptions, err = utils.PaginationFindOptions(c.DefaultQuery("page", utils.Config.APIRest.Page), c.DefaultQuery("size", utils.Config.APIRest.Size))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
				return
			}
		} else {
			findOptions = options.Find()
		}

		//Crea filtro de búsqueda para encontrar a todos los documentos
		filter := bson.D{{}}

		//Arreglo donde se almacenan los documentos encontrados
		var results []*User

		// Passing nil as the filter matches all documents in the collection
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Buscar documentos (se guardan en un cursor)
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem User
			err := cur.Decode(&elem)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
				return
			}
			elem.Password = ""
			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Cerrar el cursor
		cur.Close(context.TODO())

		//Obtener cantidad de páginas de la colección (para paginación)
		//Se salta desde página 1 hasta page-1
		if pagination {
			count, _ := collection.CountDocuments(context.TODO(), filter, options.Count().SetSkip(*findOptions.Skip))
			totalPages := utils.EstimatedTotalPages(c.DefaultQuery("size", utils.Config.APIRest.Size), count, *findOptions.Skip)
			c.Writer.Header().Set("TotalPages", totalPages)
		}

		//Retornar documentos
		c.JSON(http.StatusOK, results)
	}
}

func DeleteOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("users")

		//Obtener id del documento a eliminar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idUser"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible eliminar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Elimina el documento
		_, err = collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible eliminar el documento"})
			return
		}

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		models.CreateLog(idEx, "Eliminación de usuario", "users", c.Param("idUser"), "", "", false)

		// Eliminar todas las relaciones maps_users de este usuario
		models.DeleteAllRelMapsByUsers(id)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento eliminado exitosamente",
		})
	}
}

func UpdateOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("users")

		//De request body a User
		var elem User
		c.BindJSON(&elem)

		//Establecer fecha de actualización del documento
		//elem.UpdatedAt = time.Now()

		//Obtener id del documento a actualizar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idUser"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Valores a actualizar del documento
		update := bson.M{
			"$set": bson.M{
				"id_group":    elem.IDGroup,
				"firstname":   elem.Firstname,
				"lastname":    elem.Lastname,
				"username":    elem.Username,
				"permissions": elem.Permissions,
				//"updated_at": elem.UpdatedAt,
			},
		}

		//Actualizar documento
		_, err = collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		elem.Password = ""

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Actualización de usuario", "users", c.Param("idUser"), string(serialized), "", false)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento actualizado exitosamente",
		})
	}
}

func FindGroupByUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("users")

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idUser"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Variable donde se almacena el documento encontrado
		var result User
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		collection2 := db.Connection.Collection("groups")

		filterGroup := bson.D{primitive.E{Key: "_id", Value: result.IDGroup}}

		var resultGroup Group
		err = collection2.FindOne(context.TODO(), filterGroup).Decode(&resultGroup)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Retornar documento del grupo
		c.JSON(http.StatusOK, resultGroup)
	}
}

func FindAllMapsByUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_users")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Map

		//Obtener id del documento a buscar en los parámetros de la url
		idUser, err := primitive.ObjectIDFromHex(c.Param("idUser"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Definir pipeline
		pipeline := []bson.M{
			bson.M{"$match": bson.M{"id_user": idUser}},
			bson.M{
				"$lookup": bson.M{
					"from":         "maps",
					"localField":   "id_map",
					"foreignField": "_id",
					"as":           "mapUsers",
				},
			},
			bson.M{"$unwind": "$mapUsers"},
			bson.M{
				"$project": bson.M{
					"_id":        "$mapUsers._id",
					"name":       "$mapUsers.name",
					"desc":       "$mapUsers.desc",
					"created_at": "$mapUsers.created_at",
				},
			},
		}

		//Ejecutar pipeline
		cur, err := collection.Aggregate(context.Background(), pipeline)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Iterar sobre el cursor para guardar los documentos en el arreglo results
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem Map
			err := cur.Decode(&elem)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
				return
			}
			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Cerrar el cursor
		cur.Close(context.TODO())

		//Retornar documentos
		c.JSON(http.StatusOK, results)
	}
}

func CountUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("users")

		//Crea filtro de búsqueda para encontrar a todos los documentos
		filter := bson.D{{}}

		//Buscar documentos (se guardan en un cursor)
		count, err := collection.CountDocuments(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Retornar documentos
		c.JSON(http.StatusOK, count)

	}
}
