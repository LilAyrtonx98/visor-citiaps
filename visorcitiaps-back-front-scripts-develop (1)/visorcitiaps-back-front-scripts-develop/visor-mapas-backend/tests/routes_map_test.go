package tests

import (
	"bytes"
	"encoding/json"
	// "net/http"
	// "net/http/httptest"
	"testing"

	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/routes"
	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestMapCRUD(t *testing.T) {

	bodyPost := &models.Map{
		Name: "TEST MAP",
		Desc: "TEST DESCRIPTION"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	bodyPut := &models.Map{
		Name: "TEST MAP 2",
		Desc: "TEST DESCRIPTION 2"}

	bufPut := new(bytes.Buffer)
	json.NewEncoder(bufPut).Encode(bodyPut)

	utils.TestCRUD(
		t,
		router_routes,
		routes.InsertOneMap(),
		routes.FindAllMap(false),
		routes.FindOneMap(),
		routes.UpdateOneMap(),
		routes.DeleteOneMap(),
		"/api/test/maps",
		"/api/test/maps/:idMap",
		bufPost,
		bufPut,
		"TestMapCRUD")
}

func TestMapGetPages(t *testing.T) {
	utils.TestGET(t, router_routes, routes.FindAllMap(true), "/api/test/maps/", "TestMapGetPages")
}

func TestMapGetNoPages(t *testing.T) {
	utils.TestGET(t, router_routes, routes.FindAllMapNoPages(), "/api/test/mapsnopages/", "TestMapGetNoPages")
}

// func TestFindAllUsersByMap(t *testing.T) {
// 	testName := "TestFindAllUsersByMap"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/maps/:idMap/users", routes.FindAllUsersByMap())

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/maps/5dfcd53e248767613f05d9d1/users", nil)
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

// func TestFindAllLayersByMap(t *testing.T) {
// 	testName := "TestFindAllLayersByMap"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/maps/:idMap/layers", routes.FindAllLayersByMap())

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/maps/5dfcd53e248767613f05d9d1/layers", nil)
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

// func TestFindAllGeoprocessingsByMap(t *testing.T) {
// 	testName := "TestFindAllGeoprocessingsByMap"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/maps/:idMap/geoprocessings", routes.FindAllGeoprocessingsByMap())

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/maps/5dfcd53e248767613f05d9d1/geoprocessings", nil)
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
