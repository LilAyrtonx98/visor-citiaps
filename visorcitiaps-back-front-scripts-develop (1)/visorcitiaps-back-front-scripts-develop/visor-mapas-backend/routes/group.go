package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Group hace referencia a la estructura models.Group
type Group = models.Group

//InsertOneGroup inserta un documento en la colección groups de acuerdo al request body
func InsertOneGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("groups")

		//De request body a Group
		var elem Group
		c.BindJSON(&elem)

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

		//Actualizar ID para respone body
		elem.ID = insertResult.InsertedID.(primitive.ObjectID)

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Creación de grupo", "groups", elem.ID.Hex(), string(serialized), "", false)

		//Retornar documento
		c.JSON(http.StatusCreated, elem)
	}
}

//FindOneGroup busca un documento en la colección groups de acuerdo a su id y lo retorna en el response body
func FindOneGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("groups")

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idGroup"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Variable donde se almacena el documento encontrado
		var result Group

		//Buscar documento
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Retornar documento
		c.JSON(http.StatusOK, result)
	}
}

//FindAllGroup retorna todos los documentos de la colección grupos en el response body
func FindAllGroup(pagination bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("groups")

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
		var results []*Group

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Iterar sobre el cursor para guardar los documentos en el arreglo results
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem Group
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

//FindAllGroupNoPages retorna todos los documentos de la colección grupos en el response body (sin páginas)
func FindAllGroupNoPages() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("groups")

		//Paginación
		findOptions := options.Find()

		//Crea filtro de búsqueda para encontrar a todos los documentos
		filter := bson.D{{}}

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Group

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Iterar sobre el cursor para guardar los documentos en el arreglo results
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem Group
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

//DeleteOneGroup elimina un documento de la colección groups de acuerdo a su id
func DeleteOneGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("groups")

		//Obtener id del documento a eliminar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idGroup"))
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
		models.CreateLog(idEx, "Eliminación de grupo", "groups", c.Param("idGroup"), "", "", false)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento eliminado exitosamente",
		})
	}
}

//UpdateOneGroup actualiza un documento de la colección groups de acuerdo a su id
func UpdateOneGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("groups")

		//De request body a Group
		var elem Group
		c.BindJSON(&elem)

		//Establecer fecha de actualización del documento
		//elem.UpdatedAt = time.Now()

		//Obtener id del documento a actualizar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idGroup"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Valores a actualizar del documento
		update := bson.M{
			"$set": bson.M{
				"name": elem.Name,
				"desc": elem.Desc,
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
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Actualización de grupo", "groups", c.Param("idGroup"), string(serialized), "", false)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento actualizado exitosamente",
		})
	}
}

func FindUsersByGroup(pagination bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
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

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idGroup"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Todas las capas con id de categoria igual al consultado
		filter := bson.D{primitive.E{Key: "id_group", Value: id}}

		//Arreglo donde se almacenan los documentos encontrados
		var results []*User

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Iterar sobre el cursor para guardar los documentos en el arreglo results
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem User
			err := cur.Decode(&elem)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible actualizar el documento"})
				return
			}
			elem.Password = ""
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
		if pagination {
			count, _ := collection.CountDocuments(context.TODO(), filter, options.Count().SetSkip(*findOptions.Skip))
			totalPages := utils.EstimatedTotalPages(c.DefaultQuery("size", utils.Config.APIRest.Size), count, *findOptions.Skip)
			c.Writer.Header().Set("TotalPages", totalPages)
		}

		//Retornar documentos
		c.JSON(http.StatusOK, results)

	}
}

func CountGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("groups")

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
