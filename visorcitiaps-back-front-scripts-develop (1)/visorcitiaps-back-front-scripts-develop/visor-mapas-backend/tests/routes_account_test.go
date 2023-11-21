package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/citiaps/visor-mapas-backend/routes"
	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestFindMyUser(t *testing.T) {
	testName := "TestFindMyUser"
	// Cargar ruta al router
	router_routes.GET("/api/test/accounts/me", routes.FindMyUser())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/accounts/me", nil)
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

func TestChangeMyPassword(t *testing.T) {
	testName := "TestChangeMyPassword"
	// Cargar ruta al router
	router_routes.PUT("/api/test/accounts/password", routes.ChangeMyPassword())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("PUT", "/api/test/accounts/password", nil)
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

func TestFindMyLogs(t *testing.T) {
	testName := "TestFindMyLogs"
	// Cargar ruta al router
	router_routes.GET("/api/test/accounts/logs", routes.FindMyLogs())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/accounts/logs", nil)
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

func TestFindMyNotifications(t *testing.T) {
	testName := "TestFindMyNotifications"
	// Cargar ruta al router
	router_routes.GET("/api/test/accounts/notifications", routes.FindMyNotifications())

	// GET
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/api/test/accounts/notifications", nil)
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
