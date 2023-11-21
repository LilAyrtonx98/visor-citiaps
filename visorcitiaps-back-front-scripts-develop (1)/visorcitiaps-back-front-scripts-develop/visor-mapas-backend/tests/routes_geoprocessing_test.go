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

func TestGeoprocessingCRUD(t *testing.T) {

	bodyPost := &models.Geoprocessing{
		Name:   "TEST GEOPROCESING",
		Desc:   "TEST GEOPROCESING",
		GeoURL: "TEST URL"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	bodyPut := &models.Geoprocessing{
		Name:   "TEST GEOPROCESING 2",
		Desc:   "TEST GEOPROCESING 2",
		GeoURL: "TEST URL 2"}

	bufPut := new(bytes.Buffer)
	json.NewEncoder(bufPut).Encode(bodyPut)

	utils.TestCRUD(
		t,
		router_routes,
		routes.InsertOneGeoprocessing(),
		routes.FindAllGeoprocessing(),
		routes.FindOneGeoprocessing(),
		routes.UpdateOneGeoprocessing(),
		routes.DeleteOneGeoprocessing(),
		"/api/test/geoprocessings",
		"/api/test/geoprocessings/:idGeoprocessing",
		bufPost,
		bufPut,
		"TestGeoprocessingCRUD")
}

// func TestFindAllMapsByGeoprocessing(t *testing.T) {
// 	testName := "TestFindAllMapsByGeoprocessing"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/geoprocessings/:idGeoprocessing/maps", routes.FindAllMapsByGeoprocessing())

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/geoprocessings/5d86a5850893863547b87be0/maps", nil)
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
