package tests

import (
	"testing"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/utils"
)

/*
func TestAuthJWTMiddleware(t *testing.T) {
	// Nombre test
	test := "TestAuthJWTMiddleware"

	// Ejecutar test
	router_auth.GET("/middleware/jwt", routes.FindAllLayer()).Use(auth.AuthJWTMiddleware("admin"))

	// GET ALL
	// Crear un http request
	getRequest, err := http.NewRequest("GET", "/middleware/jwt", nil)
	if err != nil {
		t.Fatal(err)
	}
	getRequest.Header.Add("Authorization", "Bearer asjrkjsdasluqw89jas9das")
	// Registrar la respuesta
	getResponse := httptest.NewRecorder()
	// Ejecutar la solicitud
	router_auth.ServeHTTP(getResponse, getRequest)
	// Verificar resultado
	failed := getResponse.Code != http.StatusOK

	// Escribir resultado
	utils.Test(t, failed, test)

}
*/

func TestCORSMiddleware(t *testing.T) {
	// Nombre test
	test := "TestCORSMiddleware"

	// Ejecutar test
	router_auth.Use(auth.CORSMiddleware())

	//Escribir resultado
	utils.Test(t, false, test)

}
