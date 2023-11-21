package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Failed(t *testing.T, testName string) {
	t.Errorf("%v failed", testName)
}

func Success(t *testing.T, testName string) {
	t.Logf("%v success", testName)
}

func Test(t *testing.T, failed bool, testName string) {
	if failed {
		Failed(t, testName)
	} else {
		Success(t, testName)
	}
}
func TestCRUD(t *testing.T,
	router *gin.Engine,
	postHandler gin.HandlerFunc,
	getAllHandler gin.HandlerFunc,
	getOneHandler gin.HandlerFunc,
	putHandler gin.HandlerFunc,
	deleteHandler gin.HandlerFunc,
	routeName string,
	routeNameID string,
	postParams *bytes.Buffer,
	putParams *bytes.Buffer,
	testName string) {
	// Cargar ruta al router
	router.POST(routeName, postHandler)
	router.GET(routeName, getAllHandler)
	router.GET(routeNameID, getOneHandler)
	router.PUT(routeNameID, putHandler)
	router.DELETE(routeNameID, deleteHandler)

	// POST
	postRequest, err := http.NewRequest("POST", routeName, postParams)
	if err != nil {
		t.Fatal(err)
	}
	postResponse := httptest.NewRecorder()
	router.ServeHTTP(postResponse, postRequest)
	failedPost := postResponse.Code != http.StatusCreated

	// Obtener ID de objeto almacenado
	var id string
	responseData, err := ioutil.ReadAll(postResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	var raw map[string]interface{}
	json.Unmarshal([]byte(responseString), &raw)
	id = fmt.Sprintf("%v", raw["id"])

	// GET ALL
	// Crear un http request
	getAllRequest, err := http.NewRequest("GET", routeName, nil)
	if err != nil {
		t.Fatal(err)
	}
	// Registrar la respuesta
	getAllResponse := httptest.NewRecorder()
	// Ejecutar la solicitud
	router.ServeHTTP(getAllResponse, getAllRequest)
	// Verificar resultado
	failedGetAll := getAllResponse.Code != http.StatusOK

	// GET ONE
	getOneRequest, err := http.NewRequest("GET", routeName+"/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}
	getOneResponse := httptest.NewRecorder()
	router.ServeHTTP(getOneResponse, getOneRequest)
	failedGetOne := getOneResponse.Code != http.StatusOK

	// PUT
	putRequest, err := http.NewRequest("PUT", routeName+"/"+id, putParams)
	if err != nil {
		t.Fatal(err)
	}
	putResponse := httptest.NewRecorder()
	router.ServeHTTP(putResponse, putRequest)
	failedPut := putResponse.Code != http.StatusCreated

	// DELETE
	deleteRequest, err := http.NewRequest("DELETE", routeName+"/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}
	deleteResponse := httptest.NewRecorder()
	router.ServeHTTP(deleteResponse, deleteRequest)
	failedDelete := deleteResponse.Code != http.StatusCreated

	// Verificar resultado final
	failed := failedPost && failedGetAll && failedGetOne && failedPut && failedDelete
	// Escribir resultado
	Test(t, failed, testName)
}

func TestGET(t *testing.T,
	router *gin.Engine,
	getHandler gin.HandlerFunc,
	routeName string,
	testName string) {
	// Cargar ruta al router
	router.GET(routeName, getHandler)

	// GET ALL
	// Crear un http request
	getRequest, err := http.NewRequest("GET", routeName, nil)
	if err != nil {
		t.Fatal(err)
	}
	// Registrar la respuesta
	getResponse := httptest.NewRecorder()
	// Ejecutar la solicitud
	router.ServeHTTP(getResponse, getRequest)
	// Verificar resultado
	failed := getResponse.Code != http.StatusOK

	// Escribir resultado
	Test(t, failed, testName)
}
