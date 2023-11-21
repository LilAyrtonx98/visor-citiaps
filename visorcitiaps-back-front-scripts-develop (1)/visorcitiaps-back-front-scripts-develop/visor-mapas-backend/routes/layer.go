package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"strings"
	"net"
	"net/url"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Layer = models.Layer

//stringInSlice check if a string is in a slice of strings
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


//InsertOneLayer allows to save a layer in database
func InsertOneLayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access Layer Collection
		collection := db.Connection.Collection("layers")

		//Get request body Layer to elem
		var elem Layer
		if err := c.ShouldBind(&elem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		//Check IDCategory
		if elem.IDCategory.IsZero() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Check Name, Desc
		if (elem.Name == "" || elem.Desc == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Check Provider Name given
		if (elem.Provider.Name == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Provider Name convert to lower case & check supported providers Names
		elem.Provider.Name = strings.ToLower(elem.Provider.Name)
		supportedNames := []string{"file", "arcgis", "geoserver"}
		if !stringInSlice(elem.Provider.Name, supportedNames) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Provider name is not supported."})
			return
		}

		//Check Url provided
		if (elem.Provider.Url == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required."})
			return
		}
		
		//Parse Url
		parsedUrl, err := url.Parse(elem.Provider.Url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing Url."})
			return
		}
		
		//Check Protocol
		if (parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Url protocol is not supported."})
			return
		}
		elem.Provider.ParsedUrl.Protocol = parsedUrl.Scheme

		//Check Host
		host, port, _ := net.SplitHostPort(parsedUrl.Host)
		if (host == "") {
			elem.Provider.ParsedUrl.Host = parsedUrl.Host
		} else {
			elem.Provider.ParsedUrl.Host = host
			elem.Provider.ParsedUrl.Port = port
		}
		
		//Check if port is empty, complete with 80
		if (port == "") {
			elem.Provider.ParsedUrl.Port = "80"
		}
		
		//Check path
		if (parsedUrl.Path == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Url path is missing."})
			return
		}
		elem.Provider.ParsedUrl.Path = parsedUrl.Path

		//Check if url and name coincide / No es necesario porque se comprueba en el frontend (Ademas argis no necesariamente tiene nombre argis)
		// switch elem.Provider.Name {
		// 	case "arcgis":
		// 		tmp := strings.Split(parsedUrl.Path, "/")
		// 		//Warn: Url can change 'arcgis' position
		// 		if(!strings.EqualFold(tmp[1], "arcgis")) {
		// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Url does not belong to Provider Name."})
		// 			return
		// 		}
			
		// 	case "geoserver":
		// 		tmp := strings.Split(parsedUrl.Path, "/")
		// 		//Warn: Url can change 'geoserver' position
		// 		if(!strings.EqualFold(tmp[1], "geoserver")) {
		// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Url does not belong to Provider Name."})
		// 			return
		// 		}
		// 	default:
		// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Url does not belong to Any Provider Name."})
		// 		return	
		// }


		//If is geoserver, complete GeoserverData structure
		if(elem.Provider.Name == "geoserver") {
			//Access query parameters

			q := parsedUrl.Query()

			//Check service
			if q["service"] != nil {
				elem.Provider.GeoserverData.Service = q["service"][0] 
			}

			//Check version
			if q["version"] != nil {
				elem.Provider.GeoserverData.Version = q["version"][0] 
			}

			//Check Request
			if q["request"] != nil {
				elem.Provider.GeoserverData.Request = q["request"][0] 
			}

			if q["typeName"] != nil {
				//Separate Workspace and Filename
				val := strings.Split(q["typeName"][0], ":")
				elem.Provider.GeoserverData.Workspace = val[0]
				elem.Provider.GeoserverData.Filename  = val[1]
			}

			//Check maxFeatures
			if q["maxFeatures"] != nil {
				elem.Provider.GeoserverData.MaxFeatures = q["maxFeatures"][0] 
			}

			//Check OutputFormat
			if q["outputFormat"] != nil {
				elem.Provider.GeoserverData.OutputFormat = q["outputFormat"][0] 
			}
		}

		
		

		//Set CreatedAt time & UpdatedAt
		elem.CreatedAt = time.Now()
		elem.UpdatedAt = time.Now()
		
		//Insert document
		insertResult, err := collection.InsertOne(context.TODO(), elem)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Couldn't insert new document."})
			return
		}

		//Update ID for body response
		elem.ID = insertResult.InsertedID.(primitive.ObjectID)

		//Register action in log by User
		idEx, _ := auth.GetUserIDFromCookie(c)
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Creación de capa", "layers", elem.ID.Hex(), string(serialized), "", false)

		//Return statusCreated & document
		c.JSON(http.StatusCreated, elem)
	}
}

func FindOneLayer() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("layers")

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idLayer"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Variable donde se almacena el documento encontrado
		var result bson.M
		// var result Layer

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

func FindAllLayer(pagination bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("layers")

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
		var results []*Layer

		//Buscar documentos (se guardan en un cursor)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar los documentos"})
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

func DeleteOneLayer() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("layers")

		//Obtener id del documento a eliminar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idLayer"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible eliminar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Elimina el documento
		elem, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "No es posible eliminar el documento"})
			return
		}

		//Registrar acción en log de usuario
		idEx, _ := auth.GetUserIDFromCookie(c)
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Eliminación de capa", "layers", c.Param("idLayer"), string(serialized), "", false)

		// Eliminar todas las relaciones maps_layers de esta capa
		models.DeleteAllRelMapsByLayer(id)

		//Retornar mensaje confirmación
		c.JSON(http.StatusOK, gin.H{
			"message": "Documento eliminado exitosamente",
		})
	}
}

func UpdateOneLayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access Layer Collection
		collection := db.Connection.Collection("layers")

		//Get document id
		id, err := primitive.ObjectIDFromHex(c.Param("idLayer"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "Couldn't update document."})
			return
		}

		//Filter to find by ID
		filter := bson.D{primitive.E{Key: "_id", Value: id}}
		
		//Retrieve layer
		var result Layer

		//Search document
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {

			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the document."})
			return
		}

		//Get request body Layer to elem
		var elem Layer
		if err := c.ShouldBind(&elem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Now we will check if new parameters are valid
		
		//Check IDCategory
		if elem.IDCategory.IsZero() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Check Name, Desc
		if (elem.Name == "" || elem.Desc == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Check Provider Name given
		if (elem.Provider.Name == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameter is missing."})
			return
		}

		//Provider Name convert to lower case & check supported providers Names
		elem.Provider.Name = strings.ToLower(elem.Provider.Name)
		supportedNames := []string{"file", "arcgis", "geoserver"}
		if !stringInSlice(elem.Provider.Name, supportedNames) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Provider name is not supported."})
			return
		}

		//Check Url provided
		if (elem.Provider.Url == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required."})
			return
		}
		
		//Parse Url
		parsedUrl, err := url.Parse(elem.Provider.Url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing Url."})
			return
		}

		var update bson.M

		//Check if Url changes if so parse it
		if (elem.Provider.Url != result.Provider.Url) {
			//Check if URL is ok
			if (parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Url protocol is not supported."})
				return
			}
			elem.Provider.ParsedUrl.Protocol = parsedUrl.Scheme

			//Check Host
			host, port, _ := net.SplitHostPort(parsedUrl.Host)
			if (host == "") {
				elem.Provider.ParsedUrl.Host = parsedUrl.Host
			} else {
				elem.Provider.ParsedUrl.Host = host
				elem.Provider.ParsedUrl.Port = port
			}

			//Check if port is empty, complete with 80
			if (port == "") {
				elem.Provider.ParsedUrl.Port = "80"
			}

			//Check path
			if (parsedUrl.Path == "") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Url path is missing."})
				return
			}
			elem.Provider.ParsedUrl.Path = parsedUrl.Path

			//Check if url and name coincide
			switch elem.Provider.Name {
				case "arcgis":
					tmp := strings.Split(parsedUrl.Path, "/")
					//Warn: Url can change 'arcgis' position
					if(!strings.EqualFold(tmp[1], "arcgis")) {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Url does not belong to Provider Name."})
						return
					}
				
				case "geoserver":
					tmp := strings.Split(parsedUrl.Path, "/")
					//Warn: Url can change 'geoserver' position
					if(!strings.EqualFold(tmp[1], "geoserver")) {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Url does not belong to Provider Name."})
						return
					}
				default:
					c.JSON(http.StatusBadRequest, gin.H{"error": "Url does not belong to Any Provider Name."})
					return	
			}

			//If is geoserver, complete GeoserverData structure
			

			//Updated  time
			elem.UpdatedAt = time.Now()

			//Updated doc
			update = bson.M{
				"$set": bson.M{
					"id_category": elem.IDCategory,
					"name": elem.Name,
					"desc": elem.Desc,
					"provider": bson.M{
						"name": elem.Provider.Name,
						"url": elem.Provider.Url,
						"parsed_url": bson.M{
							"protocol": elem.Provider.ParsedUrl.Protocol,
							"host": elem.Provider.ParsedUrl.Host,
							"port": elem.Provider.ParsedUrl.Port,
							"path": elem.Provider.ParsedUrl.Path,
						},
					},
					"geoserverdata": bson.M{
						"service": elem.Provider.GeoserverData.Service,
						"version": elem.Provider.GeoserverData.Version,
						"request": elem.Provider.GeoserverData.Request,
						"max_features": elem.Provider.GeoserverData.MaxFeatures,
						"output_format": elem.Provider.GeoserverData.OutputFormat,
						"filename": elem.Provider.GeoserverData.Filename,
						"workspace": elem.Provider.GeoserverData.Workspace,
						"coordinates_system": elem.Provider.GeoserverData.CoordinatesSystem,
						"datastore": elem.Provider.GeoserverData.Datastore,
					},
					"updated_at": elem.UpdatedAt,
				},
			}

			//Update document with new info
			_, err = collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Couldn't update document"})
				return
			}
		}
		// else {
		// 	//New Doc
		// 	update = bson.M{
		// 		"$set": bson.M{
		// 			"id_category":    elem.IDCategory,
		// 			"name":           elem.Name,
		// 			"desc":           elem.Desc,
		// 			"provider": bson.M{
		// 				"name": elem.Provider.Name,
		// 				"url": elem.Provider.Url,
		// 			},
		// 			"updated_at": elem.UpdatedAt,
		// 		},
		// 	}

		// }

		

		//Log update action
		idEx, _ := auth.GetUserIDFromCookie(c)
		serialized, _ := json.Marshal(elem)
		models.CreateLog(idEx, "Actualización de capa", "layers", c.Param("idLayer"), string(serialized), "", false)

		//Successfully response
		c.JSON(http.StatusOK, update)
	}
}

func FindCategoryByLayer() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están las capas
		collection := db.Connection.Collection("layers")

		//Obtener id del documento a buscar en los parámetros de la url
		id, err := primitive.ObjectIDFromHex(c.Param("idLayer"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Crea filtro de búsqueda para encontrar el documento por su id
		filter := bson.D{primitive.E{Key: "_id", Value: id}}

		//Variable donde se almacena el documento encontrado
		var result Layer
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		collection2 := db.Connection.Collection("categories")

		filterGroup := bson.D{primitive.E{Key: "_id", Value: result.IDCategory}}

		var resultCategory Category
		err = collection2.FindOne(context.TODO(), filterGroup).Decode(&resultCategory)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Retornar documento del grupo
		c.JSON(http.StatusOK, resultCategory)
	}
}

func FindAllMapsByLayer() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("maps_layers")

		//Arreglo donde se almacenan los documentos encontrados
		var results []*Map

		//Obtener id del documento a buscar en los parámetros de la url
		idLayer, err := primitive.ObjectIDFromHex(c.Param("idLayer"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No es posible encontrar el documento"})
			return
		}

		//Definir pipeline
		pipeline := []bson.M{
			bson.M{"$match": bson.M{"id_layer": idLayer}},
			bson.M{
				"$lookup": bson.M{
					"from":         "maps",
					"localField":   "id_map",
					"foreignField": "_id",
					"as":           "mapLayers",
				},
			},
			bson.M{"$unwind": "$mapLayers"},
			bson.M{
				"$project": bson.M{
					"_id":        "$mapLayers._id",
					"name":       "$mapLayers.name",
					"desc":       "$mapLayers.desc",
					"created_at": "$mapLayers.created_at",
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

func CountLayer() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Acceder a la colección donde están los grupos
		collection := db.Connection.Collection("layers")

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

// func handleGeoserverXML() {
// 	resp, err := http.Get(fmt.Sprintf("%s://%s.%s:%s/%s?service=wfs&version=2.0.0&request=GetCapabilities", elem.Provider.Url.Protocol, elem.Provider.Url.Host, elem.Provider.Url.Domain, elem.Provider.Url.Port, elem.Provider.Url.Path))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Url provided does not work."})
// 		return
// 	}
	
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Url body gives an error."})
// 		return
// 	}

// 	fmt.Println(" --------------------------- ")
// 	fmt.Println(body)
// }
