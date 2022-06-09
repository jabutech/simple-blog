package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jabutech/simple-blog/post"
	"github.com/jabutech/simple-blog/router"
	"github.com/jabutech/simple-blog/util"
	"github.com/stretchr/testify/assert"
)

func createRandomPost(t *testing.T) post.Post {
	// Open connection to db
	db := util.SetupTestDb()
	// Call router with argument db
	router := router.SetupRouter(db)

	// Login for get token use LoginRandomAccount
	token := LoginRandomAccount(t)
	strToken := fmt.Sprintf("Bearer %s", token)

	// Data body with data from create account random
	dataBody := fmt.Sprintf(`{"title": "%s"}`, util.RandomString(10))
	// Create payload request
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/posts", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")
	// Added header Authorization with by inserting jwt token
	request.Header.Add("Authorization", strToken)

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
	assert.Equal(t, "Post has been created", responseBody["message"])
	// Data is not null
	assert.NotZero(t, responseBody["data"])

	var contextData = responseBody["data"].(map[string]interface{})
	// All property not empty
	assert.NotEmpty(t, contextData["id"])
	assert.NotEmpty(t, contextData["title"])

	// Map new post in object Post
	newPost := post.Post{}
	newPost.Id = contextData["id"].(string)
	newPost.Title = contextData["title"].(string)

	// Return new post
	return newPost
}

// Test Create post success
func TestCreatePostSuccess(t *testing.T) {
	createRandomPost(t)
}

// Test create post validation error
func TestCreatePostValidationError(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()
	// Call router with argument db
	router := router.SetupRouter(db)

	// Login for get token use LoginRandomAccount
	token := LoginRandomAccount(t)
	strToken := fmt.Sprintf("Bearer %s", token)

	// Data body with empty string
	dataBody := fmt.Sprintf(`{"title": ""}`)
	// Create payload request
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/posts", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")
	// Added header Authorization with by inserting jwt token
	request.Header.Add("Authorization", strToken)

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

	// Response status code must be 400 (bad request)
	assert.Equal(t, 400, response.StatusCode)
	// Response body status code must be 400 (bad request)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	// Response body status must be error
	assert.Equal(t, "error", responseBody["status"])
	// Response body message
	assert.Equal(t, "Failed create post", responseBody["message"])
	// Data is not nil
	assert.NotNil(t, responseBody["data"])
	// Error is not nil
	assert.NotNil(t, responseBody["data"].(map[string]interface{})["errors"])
}

// Test create post error title exist
func TestCreatePostTitleExistError(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()
	// Call router with argument db
	router := router.SetupRouter(db)

	// Login for get token use LoginRandomAccount
	token := LoginRandomAccount(t)
	strToken := fmt.Sprintf("Bearer %s", token)

	postExist := createRandomPost(t)
	// Data body with empty string
	dataBody := fmt.Sprintf(`{"title": "%s"}`, postExist.Title)
	// Create payload request
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/posts", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")
	// Added header Authorization with by inserting jwt token
	request.Header.Add("Authorization", strToken)

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

	// Response status code must be 400 (bad request)
	assert.Equal(t, 400, response.StatusCode)
	// Response body status code must be 400 (bad request)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	// Response body status must be error
	assert.Equal(t, "error", responseBody["status"])
	// Response body message
	assert.Equal(t, "Failed create post", responseBody["message"])
	// Data is not nil
	assert.NotNil(t, responseBody["data"])
	// Error is not nil
	assert.Equal(t, "title already exists", responseBody["data"].(map[string]interface{})["errors"])
}

// Test create post unauthorized
func TestCreatePostUnauthorized(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()
	// Call router with argument db
	router := router.SetupRouter(db)

	// Data body with data from create account random
	dataBody := fmt.Sprintf(`{"title": "%s"}`, util.RandomString(10))
	// Create payload request
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/posts", requestBody)
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

	// Response status code must be 401 (Unauthorized)
	assert.Equal(t, 401, response.StatusCode)
	// Response body status code must be 401 (Unauthorized)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	// Response body status must be error
	assert.Equal(t, "error", responseBody["status"])
	// Response body message
	assert.Equal(t, "Unauthorized", responseBody["message"])
}
