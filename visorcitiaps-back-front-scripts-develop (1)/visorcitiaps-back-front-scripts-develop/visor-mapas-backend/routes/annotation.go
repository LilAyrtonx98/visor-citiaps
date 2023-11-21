package routes

import (
	"context"
	"encoding/json"
	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type Annotation = models.Annotation

//TODO checkear que todos los parametros requeridos en para crear la anotacion son enviados en el body

/* TODO:
		- obtener las anotaciones del usuario
		- filtrar las anotaciones de grupo compartidas
		- actualizar anotacion/borrar anotacion
		- verificar que estan bien las coordenadas
 */


func FindAllAnnotationByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access Annotation Collection
		collection := db.Connection.Collection("annotations")

		//Get UserID from URL
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "Couldn't find the document."})
			return
		}

		//Pagination
		findOptions := options.Find()

		//Filter by UserID
		filter := bson.D{primitive.E{Key: "_id_user", Value: id}}

		var results []*Annotation

		//Search document
		all, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the document."})
			return
		}

		//Iterate over all to save all docs in results
		for all.Next(context.TODO()) {
			var elem Annotation
			err := all.Decode(&elem)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the document."})
				return
			}
			results = append(results, &elem)
		}

		if err := all.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		//Close all
		all.Close(context.TODO())

		//Send document
		c.JSON(http.StatusOK, results)
	}
}

func FindAllAnnotationByGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access Annotation Collection
		collection := db.Connection.Collection("annotations")

		//Get UserID from URL
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "Couldn't find the document."})
			return
		}

		//Pagination
		findOptions := options.Find()

		//Filter by UserID
		filter := bson.D{primitive.E{Key: "_id_group", Value: id}}

		var results []*Annotation

		//Search document
		all, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the document."})
			return
		}

		//Iterate over all to save all docs in results
		for all.Next(context.TODO()) {
			var elem Annotation
			err := all.Decode(&elem)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the document."})
				return
			}
			if *elem.IsShared != false {
				results = append(results, &elem)
			}
		}

		if err := all.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		//Close all
		all.Close(context.TODO())

		//Send document
		c.JSON(http.StatusOK, results)
	}
}

//InsertOneAnnotation insert a new Annotation document into his collection
func InsertOneAnnotation() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access Annotation Collection
		collection := db.Connection.Collection("annotations")

		//Get request body Annotation to elem
		var elem Annotation
		if err := c.ShouldBind(&elem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Check IDs
		if elem.IDUser.IsZero() || elem.IDMap.IsZero() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Check Shared
		f := false
		if elem.IsShared == nil {
			elem.IsShared = &f
		}

		if *elem.IsShared == true && elem.IDGroup.IsZero() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Check Location
		if elem.Location.GeoJSONType == "" || elem.Location.Coordinates == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Set CreatedAt time & UpdatedAt
		elem.CreatedAt = time.Now()
		elem.UpdatedAt = time.Now()

		//Insert document
		_, err := collection.InsertOne(context.TODO(), elem)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Couldn't insert new document."})
			return
		}

		//Register action in log by User
		idEx, _ := auth.GetUserIDFromCookie(c)
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Creación de anotación", "annotations", elem.ID.Hex(), string(serialized), "", false)

		//Return statusCreated & document
		c.JSON(http.StatusCreated, elem)
	}
}

//UpdateOneAnnotation update an Annotation document by his ID
func UpdateOneAnnotation() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access Annotation Collection
		collection := db.Connection.Collection("annotations")

		//Get request body Annotation to elem
		var elem Annotation
		if err := c.ShouldBind(&elem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Set UpdatedAt to Now
		elem.UpdatedAt = time.Now()

		//Get idAnnotation from params
		id, err := primitive.ObjectIDFromHex(c.Param("idAnnotation"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Filter to find doc by Id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		update := bson.M{
			"$set": bson.M{
				"_id_map": elem.IDMap,
				"_id_group": elem.IDGroup,
				"_is_shared": elem.IsShared,
				"text": elem.Text,
				"location": bson.M{
					"type": elem.Location.GeoJSONType,
					"coordinates": elem.Location.Coordinates,
				},
				"updated_at": elem.UpdatedAt,
			},
		}

		//Update doc
		_, err = collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Couldn't update document."})
			return
		}

		//Register action in log
		idEx,_ := auth.GetUserIDFromCookie(c)
		serialized,_ := json.Marshal(elem)
		models.CreateLog(idEx, "Actualización de anotación", "annotations", c.Param("idAnnotation"), string(serialized), "", false)

		c.JSON(http.StatusOK, gin.H{"message": "Successfully updated document."})
	}
}

func DeleteOneAnnotation() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access Annotation Collection
		collection := db.Connection.Collection("annotations")

		//Get idAnnotation from params
		id, err := primitive.ObjectIDFromHex(c.Param("idAnnotation"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Filter to find doc by Id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Delete doc
		_, err = collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Couldn't delete document."})
			return
		}

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		models.CreateLog(idEx, "Eliminación de anotación", "annotations", c.Param("idAnnotation"), "", "", false)

		c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted document."})
	}
}


