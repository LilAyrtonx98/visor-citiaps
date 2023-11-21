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

func TestGroupCRUD(t *testing.T) {

	bodyPost := &models.Group{
		Name: "TEST GROUP",
		Desc: "TEST GROUP"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	bodyPut := &models.Group{
		Name: "TEST GROUP 2",
		Desc: "TEST GROUP 2"}

	bufPut := new(bytes.Buffer)
	json.NewEncoder(bufPut).Encode(bodyPut)

	utils.TestCRUD(
		t,
		router_routes,
		routes.InsertOneGroup(),
		routes.FindAllGroup(false),
		routes.FindOneGroup(),
		routes.UpdateOneGroup(),
		routes.DeleteOneGroup(),
		"/api/test/groups",
		"/api/test/groups/:idGroup",
		bufPost,
		bufPut,
		"TestGroupCRUD")
}

func TestGroupGetNoPages(t *testing.T) {
	utils.TestGET(t, router_routes, routes.FindAllGroupNoPages(), "/api/test/groupsnopages/", "TestGroupGetNoPages")
}

func TestGroupGetPages(t *testing.T) {
	utils.TestGET(t, router_routes, routes.FindAllGroup(true), "/api/test/groups/", "TestGroupGetPages")
}

// func TestFindUsersByGroup(t *testing.T) {
// 	testName := "TestFindUsersByGroup"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/groups/:idGroup/users", routes.FindUsersByGroup(false))

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/groups/5d869c38b913f6cb5691d5af/users", nil)
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

// func TestFindUsersByGroupPages(t *testing.T) {
// 	testName := "TestFindUsersByGroupPages"
// 	// Cargar ruta al router
// 	router_routes.GET("/api/test/groups/:idGroup/users/", routes.FindUsersByGroup(true))

// 	// GET
// 	// Crear un http request
// 	getRequest, err := http.NewRequest("GET", "/api/test/groups/5d869c38b913f6cb5691d5af/users/", nil)
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
