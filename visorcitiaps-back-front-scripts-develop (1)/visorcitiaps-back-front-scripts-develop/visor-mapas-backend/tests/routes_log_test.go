package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/citiaps/visor-mapas-backend/routes"
	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestFindLogsByUserExecutor(t *testing.T) {
	testName := "TestFindLogsByUserExecutor"
	// Cargar ruta al router
	router_routes.GET("/api/test/users/:idUser/logs", routes.FindLogsByUserExecutor())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/users/5d86a5850893863547b87be0/logs", nil)
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

func TestFindLogsByUserAffected(t *testing.T) {
	testName := "TestFindLogsByUserAffected"
	// Cargar ruta al router
	router_routes.GET("/api/test/users/:idUser/logsaffected", routes.FindLogsByUserAffected())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/users/5d869d65ce92cd68db79ea40/logsaffected", nil)
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

func TestFindLogsByResourceID(t *testing.T) {
	testName := "TestFindLogsByResourceID"
	// Cargar ruta al router
	router_routes.GET("/api/test/logs/:idResource", routes.FindLogsByResourceID("layers"))

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/logs/5d86a5850893863547b87be6", nil)
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
