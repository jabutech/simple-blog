package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/auth"
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/router"
	"github.com/jabutech/simple-blog/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function setup for connection to database test
func setupTestDb() *gorm.DB {
	// Load config
	config, err := util.LoadConfig("../") // "." as location file app.env in root folder
	helper.FatalError("error load config: ", err)

	// Open connection to db
	db, err := gorm.Open(mysql.Open(config.DBSourceTest), &gorm.Config{})
	helper.FatalError("Connection to database failed!, err: ", err)

	return db
}

func setupRouter(db *gorm.DB) *gin.Engine {

	// Router
	router := router.NewRouter(db)

	return router
}

// Func for create account random
func createRandomAccount(t *testing.T, withIsAdmin bool) auth.RegisterInput {
	// Open connection to db
	db := setupTestDb()

	// Call router with argument db
	router := setupRouter(db)

	var data auth.RegisterInput
	var dataBody string

	if !withIsAdmin {
		data = auth.RegisterInput{
			Fullname: util.RandomFullname(),
			Email:    util.RandomEmail(),
			Password: "password",
		}

		dataBody = fmt.Sprintf(`{"fullname": "%s", "email": "%s", "password": "%s"}`, data.Fullname, data.Email, "password")
	} else {

		data = auth.RegisterInput{
			Fullname: util.RandomFullname(),
			Email:    util.RandomEmail(),
			Password: "password",
			IsAdmin:  true,
		}

		dataBody = fmt.Sprintf(`{"fullname": "%s", "email": "%s", "password": "%s", "is_admin": %t}`, data.Fullname, data.Email, "password", true)
	}

	// Create payload request
	requestBody := strings.NewReader(dataBody)
	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/register", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	// Create recorder
	recorder := httptest.NewRecorder()

	// Run server http
	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// Decode json
	json.Unmarshal(body, &responseBody)

	// Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// Response body status code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// Response body status must be success
	assert.Equal(t, "success", responseBody["status"])
	// Response body message
	assert.Equal(t, "You have successfully registered.", responseBody["message"])

	// Return data success register
	return data
}

// Test Register
func TestRegisterSuccessWithoutIsAdmin(t *testing.T) {
	// Var withIsAdmin value false
	withIsAdmin := false
	// Call function createRandomAccount for test create account
	createRandomAccount(t, withIsAdmin)
}

func TestRegisterSuccessWithIsAdmin(t *testing.T) {

	// Var withIsAdmin value true
	withIsAdmin := true
	// Call function createRandomAccount for test create account
	createRandomAccount(t, withIsAdmin)

}

func TestRegisterValidationError(t *testing.T) {
	// Open connection to db
	db := setupTestDb()

	// Call router with argument db
	router := setupRouter(db)

	dataBody := fmt.Sprintf(`{"fullname": "", "email": "aa", "password": "a" }`)

	// Create payload request
	requestBody := strings.NewReader(dataBody)
	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/register", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	// Create recorder
	recorder := httptest.NewRecorder()

	// Run server http
	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// Decode json
	json.Unmarshal(body, &responseBody)

	// Response status code must be 200 (success)
	assert.Equal(t, 400, response.StatusCode)
	// Response body status code must be 200 (success)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	// Response body status must be success
	assert.Equal(t, "error", responseBody["status"])
	// Response body message
	assert.Equal(t, "Registered failed.", responseBody["message"])
}

// Test Login
func TestLoginSuccess(t *testing.T) {
	// Open connection to db
	db := setupTestDb()

	// Var withIsAdmin value false
	withIsAdmin := false

	// Create account random
	account := createRandomAccount(t, withIsAdmin)
	// Call router with argument db
	router := setupRouter(db)

	// Data body with data from create account random
	dataBody := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, account.Email, account.Password)

	// Create payload request
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/login", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	// Create recorder
	recorder := httptest.NewRecorder()

	// Run server http
	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// Decode json
	json.Unmarshal(body, &responseBody)

	// Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// Response body status code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// Response body status must be success
	assert.Equal(t, "success", responseBody["status"])
	// Response body message
	assert.Equal(t, "You have Login.", responseBody["message"])
	// Data is not null
	assert.NotZero(t, responseBody["data"])
	// Property token is not nul
	assert.NotZero(t, responseBody["data"].(map[string]interface{})["token"])

}
