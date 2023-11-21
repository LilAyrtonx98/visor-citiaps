package routes

import (
	"log"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {

	//Si la función tiene un bool como parámetro: true = paginación; false = sin paginación
	//Si no, depende de cada caso para los FindAll(). Otras funciones son sin paginación.

	v1 := app.Group("/api/v1")
	{
		noAuth := v1.Group("/")
		{
			//noAuth.POST("/login", auth.LoginJWTHandler)
			noAuth.POST("/login", auth.LoginSessionHandler)
			noAuth.GET("/logout", auth.LogoutSessionHandler)

			// Para probar con TAURUS
			taurus := v1.Group("/taurus")
			taurus.GET("/layers", FindAllLayer(false)) // Sin paginación
		}

		//Rutas grupos
		groups := v1.Group("/groups")
		{
			//Proteger acciones que puedan modificar la colección
			userPermission := groups.Group("").Use(auth.AuthSessionMiddleware("users"))
			{
				userPermission.POST("", InsertOneGroup())
				userPermission.DELETE("/:idGroup", DeleteOneGroup())
				userPermission.PUT("/:idGroup", UpdateOneGroup())
			}

			//Solo requiere tener sesión iniciada
			permission := groups.Group("").Use(auth.AuthSessionMiddleware("admin"))
			{
				permission.GET("", FindAllGroup(true))   //Paginación
				permission.GET("/", FindAllGroup(false)) //Sin paginación
				permission.GET("/:idGroup", FindOneGroup())

				//Usuarios de un grupo
				permission.GET("/:idGroup/users", FindUsersByGroup(true))   //Paginación
				permission.GET("/:idGroup/users/", FindUsersByGroup(false)) //Sin paginación
			}
		}

		//Rutas usuarios
		users := v1.Group("/users")
		{
			//Proteger acciones que puedan modificar la colección
			userPermission := users.Group("").Use(auth.AuthSessionMiddleware("users"))
			{
				userPermission.POST("", InsertOneUser())
				userPermission.DELETE("/:idUser", DeleteOneUser())
				userPermission.PUT("/:idUser", UpdateOneUser())

				//Registro de acciones de un usuario
				userPermission.GET("/:idUser/logs", FindLogsByUserExecutor()) //Paginación
			}

			permission := users.Group("").Use(auth.AuthSessionMiddleware("admin"))
			{
				permission.GET("", FindAllUsers(true))   //Paginación en resultados
				permission.GET("/", FindAllUsers(false)) //Sin paginación en resultados
				permission.GET("/:idUser", FindOneUser())

				//Grupo del usuario
				permission.GET("/:idUser/groups", FindGroupByUser())

				// Mapas por usuario
				permission.GET(":idUser/maps", FindAllMapsByUser())
			}
		}

		//Rutas grupos
		layers := v1.Group("/layers")
		{
			layersPermission := layers.Group("").Use(auth.AuthSessionMiddleware("layers"))
			{
				layersPermission.POST("", InsertOneLayer())
				layersPermission.DELETE("/:idLayer", DeleteOneLayer())
				layersPermission.PUT("/:idLayer", UpdateOneLayer())
			}

			permission := layers.Group("").Use(auth.AuthSessionMiddleware("admin"))
			{
				permission.GET("", FindAllLayer(true))   //Paginación
				permission.GET("/", FindAllLayer(false)) //Sin paginación
				permission.GET("/:idLayer", FindOneLayer())

				//Categoría de la capa
				permission.GET("/:idLayer/categories", FindCategoryByLayer())

				// Mapas que tienen asignados esta capa
				permission.GET("/:idLayer/maps", FindAllMapsByLayer())
			}
		}

		//Rutas geoprocesos
		geoprocessings := v1.Group("/geoprocessings")
		{
			geoPermission := geoprocessings.Group("").Use(auth.AuthSessionMiddleware("geo"))
			{
				geoPermission.POST("", InsertOneGeoprocessing())
				geoPermission.DELETE("/:idGeoprocessing", DeleteOneGeoprocessing())
				geoPermission.PUT("/:idGeoprocessing", UpdateOneGeoprocessing())
			}

			permission := geoprocessings.Group("").Use(auth.AuthSessionMiddleware("admin"))
			{
				permission.GET("", FindAllGeoprocessing())
				permission.GET("/:idGeoprocessing", FindOneGeoprocessing())

				// Mapas que tienen asignados este geoproceso
				permission.GET("/:idGeoprocessing/maps", FindAllMapsByGeoprocessing())
			}
		}

		//Rutas categorías de capas
		categories := v1.Group("/categories")
		{
			layersPermission := categories.Group("").Use(auth.AuthSessionMiddleware("layers"))
			{
				layersPermission.POST("", InsertOneCategory())
				layersPermission.DELETE("/:idCategory", DeleteOneCategory())
				layersPermission.PUT("/:idCategory", UpdateOneCategory())
			}

			permission := categories.Group("").Use(auth.AuthSessionMiddleware("admin"))
			{
				permission.GET("", FindAllCategory(true))   //Paginación
				permission.GET("/", FindAllCategory(false)) //Sin paginación
				permission.GET("/:idCategory", FindOneCategory())

				//Capas de una categoría
				permission.GET("/:idCategory/layers", FindLayersByCategory())
			}
		}

		//Upload files
		uploads := v1.Group("/uploads")
		{
			layersPermission := uploads.Group("").Use(auth.AuthSessionMiddleware("layers"))
			{
				layersPermission.POST("/single", ReceiveSingleFile())
			}
			//uploads.POST("/singleDummy", ReceiveSingleFileDummy())
		}

		//Rutas actualización de datos usuario
		accounts := v1.Group("/accounts")
		{
			userPermission := accounts.Group("").Use(auth.AuthSessionMiddleware("users"))
			{
				userPermission.POST("/check", CheckUniqueUsername())
			}

			permission := accounts.Group("").Use(auth.AuthSessionMiddleware("admin"))
			{
				permission.GET("/me", FindMyUser())
				permission.PUT("/password", ChangeMyPassword())
				permission.GET("/logs", FindMyLogs())
				permission.GET("/notifications", FindMyNotifications())
			}
		}

		//Rutas logs (registro histórico)
		logs := v1.Group("/logs")
		{
			//Registro historico de cambios asociados a un recurso
			userPermission := logs.Group("").Use(auth.AuthSessionMiddleware("users"))
			{
				userPermission.GET("/users/:idResource", FindLogsByResourceID("users"))
				userPermission.GET("/groups/:idResource", FindLogsByResourceID("groups"))
			}
			layersPermission := logs.Group("").Use(auth.AuthSessionMiddleware("layers"))
			{
				layersPermission.GET("/layers/:idResource", FindLogsByResourceID("layers"))
				layersPermission.GET("/categories/:idResource", FindLogsByResourceID("categories"))
			}
			geoPermission := logs.Group("").Use(auth.AuthSessionMiddleware("geo"))
			{
				geoPermission.GET("/geoprocessings/:idResource", FindLogsByResourceID("geoprocessings"))
			}
		}

		//Rutas count
		counts := v1.Group("/counts")
		{
			permission := counts.Group("").Use(auth.AuthSessionMiddleware("none"))
			{
				permission.GET("/maps", CountMap())
				permission.GET("/layers", CountLayer())
				permission.GET("/categories", CountCategory())
				permission.GET("/users", CountUser())
				permission.GET("/groups", CountGroup())
				permission.GET("/geoprocessings", CountGeoprocessing())
			}
		}

		//Upload files
		uploads2 := v1.Group("/uploads2") //Solo para demostrar como funciona la subida de archivos a GeoServer
		{
			uploads2.POST("/singleDummy", ReceiveSingleFileDummy())
		}

		//Rutas visor de mapas
		maps := v1.Group("/maps")
		{
			// TODO: revisar y decidir si seguir usando permiso de capas o agregar un nuevo permiso para mapas
			layersPermission := maps.Group("").Use(auth.AuthSessionMiddleware("maps"))
			{
				layersPermission.POST("", InsertOneMap())
				layersPermission.DELETE("/:idMap", DeleteOneMap())
				layersPermission.PUT("/:idMap", UpdateOneMap())
			}

			permission := maps.Group("").Use(auth.AuthSessionMiddleware("admin"))
			{
				permission.GET("", FindAllMap(true))   //Paginación
				permission.GET("/", FindAllMap(false)) //Sin paginación
				permission.GET("/:idMap", FindOneMap())

				permission.GET(":idMap/users", FindAllUsersByMap())
				permission.GET(":idMap/layers", FindAllLayersByMap())
				permission.GET(":idMap/geoprocessings", FindAllGeoprocessingsByMap())
				permission.GET(":idMap/layers/categories", FindAllLayersSortedByCategorysByMap())

				// Asignar capa a mapa
				permission.POST(":idMap/layers/:idLayer", CreateRelMapLayer())
				
				// No asignar capa a mapa
				permission.DELETE(":idMap/layers/:idLayer", DeleteRelMapLayer())

				//Asignar geoproceso a mapa
				permission.POST(":idMap/geoprocessings/:idGeoprocessing", CreateRelMapGeo())
				//No asignar geoproceso a mapa
				permission.DELETE(":idMap/geoprocessings/:idGeoprocessing", DeleteRelMapGeo())

				//Asignar usuario a mapa
				permission.POST(":idMap/users/:idUser", CreateRelMapUser())
				//No asignar usuario a mapa
				permission.DELETE(":idMap/users/:idUser", DeleteRelMapUser())
			}
		}

		//Rutas visor de mapas
		visor := v1.Group("/visor")
		{
			visorPermission := visor.Group("").Use(auth.AuthSessionMiddleware("visor"))
			{
				// Mapas disponibles para el usuario
				visorPermission.GET("/maps", FindAllMapsByUserCookieID())
				// Capas ordenadas por categoría dentro de un mapa disponibles para el usuario
				visorPermission.GET("/maps/:idMap/layers", FindAllLayersSortedByCategoryInOneMapByUserCookieID())
				// Geoprocesos dentro de un mapa disponibles para el usuario
				visorPermission.GET("/maps/:idMap/geoprocessings", FindAllGeoprocessingsInOneMapByUserCookieID())
			}
		}

		//Annotation routes
		annotation := v1.Group("/annotations")
		{
			annotationPermission := annotation.Group("").Use(auth.AuthSessionMiddleware("annotation"))
			{
				//Get annotations by UserID
				annotationPermission.GET("/:id", FindAllAnnotationByUser())
				//Get annotations by GroupID
				annotationPermission.GET("/:id/group", FindAllAnnotationByGroup())
				//Insert
				annotationPermission.POST("", InsertOneAnnotation())
				//Update
				annotationPermission.PUT("/:idAnnotation", UpdateOneAnnotation())
				//Delete
				annotationPermission.DELETE("/:idAnnotation", DeleteOneAnnotation())
			}
		}
	}
	log.Println("Rutas cargadas")
}
