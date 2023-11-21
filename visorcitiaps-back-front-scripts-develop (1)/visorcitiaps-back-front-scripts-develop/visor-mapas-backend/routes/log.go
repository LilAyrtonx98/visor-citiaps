package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Group hace referencia a la estructura models.Group
type Log = models.Log

func FindLogsByUserExecutor() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("logs")

		//Paginación
		findOptions, err := utils.PaginationFindOptions(c.DefaultQuery("page", utils.Config.APIRest.Page), c.DefaultQuery("size", utils.Config.APIRest.Size))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Todas las capas con id de categoria igual al consultado
		filter := bson.M{"id_user_executor": c.Param("idUser")}

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

func FindLogsByUserAffected() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("logs")

		//Paginación
		findOptions, err := utils.PaginationFindOptions(c.DefaultQuery("page", utils.Config.APIRest.Page), c.DefaultQuery("size", utils.Config.APIRest.Size))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Todas las capas con id de categoria igual al consultado
		filter := bson.M{
			"id_user_affected":     c.Param("idUser"),
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

func FindLogsByResourceID(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("logs")

		//Paginación
		findOptions, err := utils.PaginationFindOptions(c.DefaultQuery("page", utils.Config.APIRest.Page), c.DefaultQuery("size", utils.Config.APIRest.Size))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Todas las capas con id de categoria igual al consultado
		filter := bson.M{
			"collection":  resource,
			"id_resource": c.Param("idResource"),
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
