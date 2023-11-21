package tests

/*
func TestLoginJWTHandler(t *testing.T) {
	testName := "TestLoginJWTHandler"

	router_auth.POST("/api/test/login/jwt", auth.LoginJWTHandler)

	bodyPost := &auth.Credentials{
		Username: "user1@testing.cl",
		Password: "testing"}

	bufPost := new(bytes.Buffer)
	json.NewEncoder(bufPost).Encode(bodyPost)

	// POST
	postRequest, err := http.NewRequest("POST", "/api/test/login/jwt", bufPost)
	if err != nil {
		t.Fatal(err)
	}
	postResponse := httptest.NewRecorder()
	router_auth.ServeHTTP(postResponse, postRequest)

	failed := postResponse.Code != 200

	utils.Test(t, failed, testName)
}
*/
