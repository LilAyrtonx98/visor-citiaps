package main

import (
	"log"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/routes"
	"github.com/citiaps/visor-mapas-backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	//Cargar configuraciones
	utils.LoadConfig("./config/config.yml")

	// Mensaje
	log.Println("Aplicación disponible en " + utils.Config.Server.Host + ":" + utils.Config.Server.Port)
	

	// Logs se escriben en archivo
	utils.InitLogger()

	// Conectar con DB
	db.Setup()

	// Iniciar auth session
	auth.InitAuthSession()

	// Iniciar auth JWT
	//auth.InitJWTSession()

	// Iniciar SMTP
	utils.InitSMTPServer()

	// Aplicación GIN
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	// Configurar CORS
	app.Use(auth.CORSMiddleware())

	// Cargar rutas
	routes.Setup(app)

	// Mensaje (logfile.log)
	log.Println("Aplicación disponible en " + utils.Config.Server.Host + ":" + utils.Config.Server.Port)
	log.Println("Dominio cargado ")

	log.Println("Dominio cargado " + utils.Config.CORS.Origin)

	// Ejecutar aplicación
	app.Run(":" + utils.Config.Server.Port) // listen and serve on 0.0.0.0:<port>
}
