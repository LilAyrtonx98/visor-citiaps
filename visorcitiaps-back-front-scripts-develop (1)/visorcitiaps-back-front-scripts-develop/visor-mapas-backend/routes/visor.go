package routes

import (
	"context"
	"net/http"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************
	Utilizadas por los usuarios en el visor de mapas
***************/

func FindAllMapsByUserCookieID() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_users")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Map

		idString, ok := auth.GetUserIDFromCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Obtener id del documento a buscar en el id string
		idUser, err := primitive.ObjectIDFromHex(idString)
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
					"imgurl":     "$mapUsers.imgurl",
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

func FindAllLayersSortedByCategoryInOneMapByUserCookieID() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Obtener id del usuario que se encuentra en la cookie
		idString, ok := auth.GetUserIDFromCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Obtener id del documento a buscar en el id string
		idUser, err := primitive.ObjectIDFromHex(idString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Obtener id del documento a buscar en los parámetros de la url
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		collection0 := db.Connection.Collection("maps_users")

		filterCheck := bson.D{primitive.E{Key: "id_map", Value: idMap}, primitive.E{Key: "id_user", Value: idUser}}

		var result RelationshipMapUser

		// Verificar que el usuario tiene asignado el mapa
		err = collection0.FindOne(context.TODO(), filterCheck).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No tienes permiso para acceder a este mapa"})
			return
		}

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_layers")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*CategoryWithLayers

		//Definir pipeline
		pipeline := []bson.M{
			bson.M{"$match": bson.M{"id_map": idMap}},
			bson.M{
				"$lookup": bson.M{
					"from":         "layers",
					"localField":   "id_layer",
					"foreignField": "_id",
					"as":           "layersMap",
				},
			},
			bson.M{"$unwind": "$layersMap"},

			bson.M{"$group": bson.M{"_id": "$layersMap.id_category", "layers": bson.M{"$push": "$layersMap"}}},

			bson.M{
				"$lookup": bson.M{
					"from":         "categories",
					"localField":   "_id",
					"foreignField": "_id",
					"as":           "categories",
				},
			},

			bson.M{"$unwind": "$categories"},

			bson.M{
				"$project": bson.M{
					"_id":        "$_id",
					"name":       "$categories.name",
					"desc":       "$categories.desc",
					"layers":     "$layers",
					"created_at": "$categories.created_at",
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
			var elem CategoryWithLayers
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
		c.JSON(200, results)
	}
}

func FindAllGeoprocessingsInOneMapByUserCookieID() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Obtener id del usuario que se encuentra en la cookie
		idString, ok := auth.GetUserIDFromCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Obtener id del documento a buscar en el id string
		idUser, err := primitive.ObjectIDFromHex(idString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Obtener id del documento a buscar en los parámetros de la url
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		collection0 := db.Connection.Collection("maps_users")

		filterCheck := bson.D{primitive.E{Key: "id_map", Value: idMap}, primitive.E{Key: "id_user", Value: idUser}}

		var result RelationshipMapUser

		// Verificar que el usuario tiene asignado el mapa
		err = collection0.FindOne(context.TODO(), filterCheck).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No tienes permiso para acceder a este mapa"})
			return
		}

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_geoprocessings")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Geoprocessing

		//Definir pipeline
		pipeline := []bson.M{
			bson.M{"$match": bson.M{"id_map": idMap}},
			bson.M{
				"$lookup": bson.M{
					"from":         "geoprocessings",
					"localField":   "id_geoprocessing",
					"foreignField": "_id",
					"as":           "geoprocessingsMap",
				},
			},
			bson.M{"$unwind": "$geoprocessingsMap"},
			bson.M{
				"$project": bson.M{
					"_id":        "$geoprocessingsMap._id",
					"name":       "$geoprocessingsMap.name",
					"desc":       "$geoprocessingsMap.desc",
					"geo_url":    "$geoprocessingsMap.geo_url",
					"created_at": "$geoprocessingsMap.created_at",
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
			var elem Geoprocessing
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
