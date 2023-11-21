package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RelationshipMapUser = models.RelationshipMapUser

type RelationshipMapGeoprocessing = models.RelationshipMapGeoprocessing

type RelationshipMapLayer = models.RelationshipMapLayer

func CreateRelMapUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las relaciones
		collection := db.Connection.Collection("maps_users")

		//De request body a Relationship
		var elem RelationshipMapUser

		//Establecer fecha de creación del documento
		elem.CreatedAt = time.Now()

		//ID del usuario al cual se le asigna la capa
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}
		elem.IDMap = idMap

		//ID de la capa que se le asigana al usuario
		idUser, err := primitive.ObjectIDFromHex(c.Param("idUser"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}
		elem.IDUser = idUser

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
		models.CreateLog(idEx, "Asignación de geoproceso", "maps_users", elem.IDUser.Hex(), "", elem.IDMap.Hex(), true)

		//Notificar por correo (asíncrono)
		models.NotifyUserAssigedResource(idMap, idUser, "mapa")

		//Retornar elemento
		c.JSON(http.StatusCreated, elem)
	}
}

func DeleteRelMapUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("maps_users")

		//ID del mapa al cual se le asigna la capa
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//ID de la capa que se le asigana al usuario
		idUser, err := primitive.ObjectIDFromHex(c.Param("idUser"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{
			primitive.E{
				Key: "id_map", Value: idMap,
			},
			primitive.E{
				Key: "id_user", Value: idUser,
			},
		}

		//Elimina el documento
		_, err = collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible eliminar el documento"})
			return
		}

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		models.CreateLog(idEx, "No asignación de geoproceso", "maps_users", c.Param("idUser"), "", c.Param("idMap"), true)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento eliminado exitosamente",
		})
	}
}

func CreateRelMapGeo() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las relaciones
		collection := db.Connection.Collection("maps_geoprocessings")

		//De request body a Relationship
		var elem RelationshipMapGeoprocessing

		//Establecer fecha de creación del documento
		elem.CreatedAt = time.Now()

		//ID del usuario al cual se le asigna la capa
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}
		elem.IDMap = idMap

		//ID de la capa que se le asigana al usuario
		idGeoprocessing, err := primitive.ObjectIDFromHex(c.Param("idGeoprocessing"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}
		elem.IDGeoprocessing = idGeoprocessing

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
		models.CreateLog(idEx, "Asignación de geoproceso", "maps_geoprocessings", elem.IDGeoprocessing.Hex(), "", elem.IDMap.Hex(), true)

		//Notificar por correo (asíncrono)
		models.NotifyUserAssigedResource(idMap, idGeoprocessing, "geoprocessing")

		//Retornar elemento
		c.JSON(http.StatusCreated, elem)
	}
}

func DeleteRelMapGeo() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("maps_geoprocessings")

		//ID del mapa al cual se le asigna la capa
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//ID de la capa que se le asigana al usuario
		idGeoprocessing, err := primitive.ObjectIDFromHex(c.Param("idGeoprocessing"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{
			primitive.E{
				Key: "id_map", Value: idMap,
			},
			primitive.E{
				Key: "id_geoprocessing", Value: idGeoprocessing,
			},
		}

		//Elimina el documento
		_, err = collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible eliminar el documento"})
			return
		}

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		models.CreateLog(idEx, "No asignación de geoproceso", "maps_geoprocessings", c.Param("idGeoprocessing"), "", c.Param("idMap"), true)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento eliminado exitosamente",
		})
	}
}

func CreateRelMapLayer() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las relaciones
		collection := db.Connection.Collection("maps_layers")

		//De request body a Relationship
		var elem RelationshipMapLayer

		//Establecer fecha de creación del documento
		elem.CreatedAt = time.Now()

		//ID del usuario al cual se le asigna la capa
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}
		elem.IDMap = idMap

		//ID de la capa que se le asigana al usuario
		idLayer, err := primitive.ObjectIDFromHex(c.Param("idLayer"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}
		elem.IDLayer = idLayer

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
		models.CreateLog(idEx, "Asignación de capa", "maps_layers", elem.IDLayer.Hex(), "", elem.IDMap.Hex(), true)

		//Notificar por correo (asíncrono)
		models.NotifyUserAssigedResource(idMap, idLayer, "layer")

		//Retornar elemento
		c.JSON(http.StatusCreated, elem)
	}
}

func DeleteRelMapLayer() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los usuarios
		collection := db.Connection.Collection("maps_layers")

		//ID del usuario al cual se le asigna la capa
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//ID de la capa que se le asigana al usuario
		idLayer, err := primitive.ObjectIDFromHex(c.Param("idLayer"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible insertar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{
			primitive.E{
				Key: "id_map", Value: idMap,
			},
			primitive.E{
				Key: "id_layer", Value: idLayer,
			},
		}

		//Elimina el documento
		_, err = collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible eliminar el documento"})
			return
		}

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		models.CreateLog(idEx, "No asignación de capa", "maps_layers", c.Param("idLayer"), "", c.Param("idMap"), true)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento eliminado exitosamente",
		})
	}
}
