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

func TestCategoryCRUD(t *testing.T) {

	bodyPost := &models.Category{
		Name: "TEST CATEGORY",
		Desc: "TEST DESCRIPTION"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	bodyPut := &models.Category{
		Name: "TEST CATEGORY 2",
		Desc: "TEST DESCRIPTION 2"}

	bufPut := new(bytes.Buffer)
	json.NewEncoder(bufPut).Encode(bodyPut)

	utils.TestCRUD(
		t,
		router_routes,
		routes.InsertOneCategory(),
		routes.FindAllCategory(false),
		routes.FindOneCategory(),
		routes.UpdateOneCategory(),
		routes.DeleteOneCategory(),
		"/api/test/categories",
		"/api/test/categories/:idCategory",
		bufPost,
		bufPut,
		"TestCategoryCRUD")
}

func TestCategoryGetPages(t *testing.T) {
	utils.TestGET(t, router_routes, routes.FindAllCategory(true), "/api/test/categories/", "TestCategoryGetPages")
}

func TestCategoryGetNoPages(t *testing.T) {
	utils.TestGET(t, router_routes, routes.FindAllCategoryNoPages(), "/api/test/categoriesnopages/", "TestCategoryGetNoPages")
}

// func TestFindLayersByCategory(t *testing.T) {
// 	testName := "TestFindLayersByCategory"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/categories/:idCategory/layers", routes.FindLayersByCategory())

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/categories/5dfcd53e248767613f05d9d1/layers", nil)
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
