package tests

import (
	// "net/http"
	// "net/http/httptest"
	// "testing"
	// "gitlab.com/cesar.kreep/visor-mapas-backend/routes"
	// "gitlab.com/cesar.kreep/visor-mapas-backend/utils"
)

// func TestCreateRelMapUser(t *testing.T) {
// 	testName := "TestCreateRelMapUser"

// 	router_routes.POST("/api/test/relationships/maps/:idMap/users/:idUser", routes.CreateRelMapUser())

// 	// POST
// 	postRequest, err := http.NewRequest("POST", "/api/test/relationships/maps/5d86a5850893863547b87be0/users/5d869d65ce92cd68db79ea40", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	postResponse := httptest.NewRecorder()
// 	router_routes.ServeHTTP(postResponse, postRequest)

// 	failed := postResponse.Code != http.StatusCreated

// 	utils.Test(t, failed, testName)
// }

// func TestDeleteRelMapUser(t *testing.T) {
// 	testName := "TestDeleteRelMapUser"

// 	router_routes.DELETE("/api/test/relationships/maps/:idMap/users/:idUser", routes.DeleteRelMapUser())

// 	// DELETE
// 	deleteRequest, err := http.NewRequest("DELETE", "/api/test/relationships/maps/5d86a5850893863547b87be0/users/5d86a5850893863547b87be1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	deleteResponse := httptest.NewRecorder()
// 	router_routes.ServeHTTP(deleteResponse, deleteRequest)

// 	// Verificar resultado final
// 	failed := deleteResponse.Code != http.StatusOK

// 	utils.Test(t, failed, testName)
// }

// func TestCreateRelMapGeo(t *testing.T) {
// 	testName := "TestCreateRelMapGeo"

// 	router_routes.POST("/api/test/relationships/maps/:idMap/geoprocessings/:idGeoprocessing", routes.CreateRelMapGeo())

// 	// POST
// 	postRequest, err := http.NewRequest("POST", "/api/test/relationships/maps/5d86a5850893863547b87be0/geoprocessings/5d86a5850893863547b87be2", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	postResponse := httptest.NewRecorder()
// 	router_routes.ServeHTTP(postResponse, postRequest)

// 	failed := postResponse.Code != http.StatusCreated

// 	utils.Test(t, failed, testName)
// }

// func TestDeleteRelMapGeo(t *testing.T) {
// 	testName := "TestDeleteRelMapGeo"

// 	router_routes.DELETE("/api/test/relationships/maps/:idMap/geoprocessings/:idGeoprocessing", routes.DeleteRelMapGeo())

// 	// DELETE
// 	deleteRequest, err := http.NewRequest("DELETE", "/api/test/relationships/maps/5d86a5850893863547b87be0/geoprocessings/5d86a5850893863547b87be2", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	deleteResponse := httptest.NewRecorder()
// 	router_routes.ServeHTTP(deleteResponse, deleteRequest)

// 	// Verificar resultado final
// 	failed := deleteResponse.Code != http.StatusOK

// 	utils.Test(t, failed, testName)

// }

// func TestCreateRelMapLayer(t *testing.T) {
// 	testName := "TestCreateRelMapLayer"

// 	router_routes.POST("/api/test/relationships/maps/:idMap/layers/:idLayer", routes.CreateRelMapLayer())

// 	// POST
// 	postRequest, err := http.NewRequest("POST", "/api/test/relationships/maps/5d86a5850893863547b87be0/layers/5d86a5850893863547b87be3", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	postResponse := httptest.NewRecorder()
// 	router_routes.ServeHTTP(postResponse, postRequest)

// 	failed := postResponse.Code != http.StatusCreated

// 	utils.Test(t, failed, testName)
// }

// func TestDeleteRelMapLayer(t *testing.T) {
// 	testName := "TestDeleteRelMapLayer"

// 	router_routes.DELETE("/api/test/relationships/maps/:idMap/layers/:idLayer", routes.DeleteRelMapLayer())

// 	DELETE
// 	deleteRequest, err := http.NewRequest("DELETE", "/api/test/relationships/maps/5d86a5850893863547b87be0/layers/5d86a5850893863547b87be3", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	deleteResponse := httptest.NewRecorder()
// 	router_routes.ServeHTTP(deleteResponse, deleteRequest)

// 	Verificar resultado final
// 	failed := deleteResponse.Code != http.StatusOK

// 	utils.Test(t, failed, testName)

// }
