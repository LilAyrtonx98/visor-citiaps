package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestLoginSessionHandler(t *testing.T) {
	testName := "TestLoginSessionHandler"

	router_auth.POST("/api/test/login/session", auth.LoginSessionHandler)

	bodyPost := &auth.Credentials{
		Username: "admin@visor.cl",
		Password: "holahola"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	// POST
	postRequest, err := http.NewRequest("POST", "/api/test/login/session", bufPost)
	if err != nil {
		t.Fatal(err)
	}
	postResponse := httptest.NewRecorder()
	router_auth.ServeHTTP(postResponse, postRequest)

	failed := postResponse.Code != http.StatusOK

	utils.Test(t, failed, testName)
}

func TestLogoutSessionHandler(t *testing.T) {
	testName := "TestLogoutSessionHandler"

	router_auth.GET("/api/test/login/session/logout", auth.LogoutSessionHandler)

	// POST
	postRequest, err := http.NewRequest("GET", "/api/test/login/session/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	postResponse := httptest.NewRecorder()
	router_auth.ServeHTTP(postResponse, postRequest)

	failed := postResponse.Code != http.StatusCreated && postResponse.Code != http.StatusBadRequest

	utils.Test(t, failed, testName)
}
