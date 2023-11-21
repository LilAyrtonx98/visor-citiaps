package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/routes"
	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestLayerCRUD(t *testing.T) {

	bodyPost := &models.Layer{
		Name: "TEST LAYER",
		Desc: "TEST LAYER"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	bodyPut := &models.Layer{
		Name: "TEST LAYER 2",
		Desc: "TEST LAYER 2"}

	bufPut := new(bytes.Buffer)
	json.NewEncoder(bufPut).Encode(bodyPut)

	utils.TestCRUD(
		t,
		router_routes,
		routes.InsertOneLayer(),
		routes.FindAllLayer(false),
		routes.FindOneLayer(),
		routes.UpdateOneLayer(),
		routes.DeleteOneLayer(),
		"/api/test/layers",
		"/api/test/layers/:idLayer",
		bufPost,
		bufPut,
		"TestLayerCRUD")
}

func TestFindCategoryByLayer(t *testing.T) {
	testName := "TestFindCategoryByLayer"
	// Cargar ruta al router
	router_routes.GET("/api/test/layers/:idLayer/categories", routes.FindCategoryByLayer())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/layers/5d86a5850893863547b87be6/categories", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Registrar la respuesta
	getResponse := httptest.NewRecorder()
	// Ejecutar la solicitud
	router_routes.ServeHTTP(getResponse, getRequest)
	// Verificar resultado
	failed := err != nil

	// Escribir resultado
	utils.Test(t, failed, testName)
}

// func TestFindAllMapsByLayer(t *testing.T) {
// 	testName := "TestFindAllMapsByLayer"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/layers/:idLayer/maps", routes.FindAllMapsByLayer())

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/layers/5d86a5850893863547b87be6/maps", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// Registrar la respuesta
// 	getResponse := httptest.NewRecorder()
// 	// Ejecutar la solicitud
// 	router_routes.ServeHTTP(getResponse, getRequest)
// 	// Verificar resultado
// 	failed := getResponse.Code != http.StatusOK

// 	// Escribir resultado
// 	utils.Test(t, failed, testName)
// }
