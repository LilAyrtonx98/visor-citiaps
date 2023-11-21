package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/citiaps/visor-mapas-backend/routes"
	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestRoutesSetup(t *testing.T) {
	app := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.Setup(app)
	// Escribir resultado
	utils.Test(t, false, "TestRoutesSetup")
}
