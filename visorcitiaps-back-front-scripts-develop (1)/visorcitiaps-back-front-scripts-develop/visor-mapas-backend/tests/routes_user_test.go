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

func TestUserCRUD(t *testing.T) {

	bodyPost := &models.User{
		Firstname: "TEST USER",
		Lastname:  "TEST USER"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	bodyPut := &models.User{
		Firstname: "TEST USER 2",
		Lastname:  "TEST USER 2"}

	bufPut := new(bytes.Buffer)
	json.NewEncoder(bufPut).Encode(bodyPut)

	utils.TestCRUD(
		t,
		router_routes,
		routes.InsertOneUser(),
		routes.FindAllUsers(false),
		routes.FindOneUser(),
		routes.UpdateOneUser(),
		routes.DeleteOneUser(),
		"/api/test/users",
		"/api/test/users/:idUser",
		bufPost,
		bufPut,
		"TestUserCRUD")
}

func TestUserGetPages(t *testing.T) {
	utils.TestGET(t, router_routes, routes.FindAllUsers(true), "/api/test/users/", "TestUserGetPages")
}

func TestFindGroupByUser(t *testing.T) {
	testName := "TestFindGroupByUser"
	// Cargar ruta al router
	router_routes.GET("/api/test/users/:idUser/groups", routes.FindGroupByUser())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/users/5d869d65ce92cd68db79ea40/groups", nil)
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

func TestFindAllMapsByUser(t *testing.T) {
	testName := "TestFindAllMapsByUser"
	// Cargar ruta al router
	router_routes.GET("/api/test/users/:idUser/maps", routes.FindAllMapsByUser())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/users/5d869d65ce92cd68db79ea40/maps", nil)
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
