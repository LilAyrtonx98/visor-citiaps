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

//Group hace referencia a la estructura models.Map
type Map = models.Map

//InsertOneGroup inserta un documento en la colección groups de acuerdo al request body
func InsertOneMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps")

		//De request body a Role
		var elem Map
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
		models.CreateLog(idEx, "Creación de mapa", "maps", elem.ID.Hex(), string(serialized), "", false)

		c.JSON(http.StatusCreated, elem)
	}
}

//FindOneGroup busca un documento en la colección groups de acuerdo a su id y lo retorna en el response body
func FindOneMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps")

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Variable donde se almacena el documento encontrado
		var result Map

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
func FindAllMap(pagination bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps")

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
		var results []*Map

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
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

//FindAllMapNoPage retorna todos los documentos de la colección grupos en el response body
func FindAllMapNoPages() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps")

		//Paginación
		findOptions := options.Find()

		//Crea filtro de búsqueda para encontrar a todos los documentos
		filter := bson.D{{}}

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Map

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
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

//DeleteOneGroup elimina un documento de la colección groups de acuerdo a su id
func DeleteOneMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps")

		//Obtener id del documento a eliminar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idMap"))
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
		models.CreateLog(idEx, "Eliminación de mapa", "maps", c.Param("idMap"), "", "", false)

		// Eliminar todas las relaciones maps_users, maps_geoprocessings y maps_layers de este mapa
		models.DeleteAllRelByMap(id)

		c.JSON(http.StatusOK, gin.H{
			"message": "Documento eliminado exitosamente",
		})
	}
}

//UpdateOneGroup actualiza un documento de la colección groups de acuerdo a su id
func UpdateOneMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps")

		//De request body a Group
		var elem Map
		c.BindJSON(&elem)

		//Establecer fecha de actualización del documento
		//elem.UpdatedAt = time.Now()

		//Obtener id del documento a actualizar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible actualizar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Valores a actualizar del documento
		update := bson.M{
			"$set": bson.M{
				"name":   elem.Name,
				"desc":   elem.Desc,
				"imgurl": elem.ImgURL,
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
		models.CreateLog(idEx, "Actualización de mapa", "maps", c.Param("idMap"), string(serialized), "", false)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento actualizado exitosamente",
		})
	}
}

func FindLayersByMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("layers")

		//Paginación
		findOptions, err := utils.PaginationFindOptions(c.DefaultQuery("page", utils.Config.APIRest.Page), c.DefaultQuery("size", utils.Config.APIRest.Size))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
			return
		}

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Todas las capas con id de categoria igual al consultado
		filter := bson.D{primitive.E{Key: "id_category", Value: id}}

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Layer

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Iterar sobre el cursor para guardar los documentos en el arreglo results
		for cur.Next(context.TODO()) {

			//Variable donde se almacena el documento del cursor
			var elem Layer
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

func CountMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps")

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

func FindAllUsersByMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_users")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*User

		//Obtener id del documento a buscar en los parámetros de la url
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Definir pipeline
		pipeline := []bson.M{
			bson.M{"$match": bson.M{"id_map": idMap}},
			bson.M{
				"$lookup": bson.M{
					"from":         "users",
					"localField":   "id_user",
					"foreignField": "_id",
					"as":           "usersMap",
				},
			},
			bson.M{"$unwind": "$usersMap"},
			bson.M{
				"$project": bson.M{
					"_id":        "$usersMap._id",
					"id_group":   "$usersMap.id_group",
					"firstname":  "$usersMap.firstname",
					"lastname":   "$usersMap.lastname",
					"username":   "$usersMap.username",
					"password":   "$usersMap.password",
					"created_at": "$usersMap.created_at",
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
			var elem User
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

func FindAllLayersByMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_layers")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Layer

		//Obtener id del documento a buscar en los parámetros de la url
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

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
			bson.M{
				"$project": bson.M{
					"_id":         "$layersMap._id",
					"id_category": "$layersMap.id_category",
					"name":        "$layersMap.name",
					"desc":        "$layersMap.desc",
					"layer_url":   "$layersMap.layer_url",
					"created_at":  "$layersMap.created_at",
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
			var elem Layer
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

func FindAllGeoprocessingsByMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_geoprocessings")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Geoprocessing

		//Obtener id del documento a buscar en los parámetros de la url
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

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

func FindAllLayersSortedByCategorysByMap() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_layers")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*CategoryWithLayers

		//Obtener id del documento a buscar en los parámetros de la url
		idMap, err := primitive.ObjectIDFromHex(c.Param("idMap"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

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
