package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/post"
	"github.com/jabutech/simple-blog/router"
	"github.com/jabutech/simple-blog/user"
	"github.com/jabutech/simple-blog/util"
	"github.com/stretchr/testify/assert"
)

func createRandomPost(t *testing.T, isAdminTrue bool) (post.Post, string) {
	// Open connection to db
	db := util.SetupTestDb()
	// Call router with argument db
	router := router.SetupRouter(db)

	// Login for get token use LoginRandomAccount, with argument true set as admin `true`
	token := LoginRandomAccount(t, isAdminTrue)
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

	// Return newPost, and token used to create this post
	return newPost, strToken
}

// Test Create post success
func TestCreatePostSuccess(t *testing.T) {
	createRandomPost(t, true)
}

// Test create post validation error
func TestCreatePostValidationError(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()
	// Call router with argument db
	router := router.SetupRouter(db)

	// Login for get token use LoginRandomAccount, with argument true set as admin `true`
	token := LoginRandomAccount(t, true)
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

	// Login for get token use LoginRandomAccount, with argument true set as admin `true`
	token := LoginRandomAccount(t, true)
	strToken := fmt.Sprintf("Bearer %s", token)

	postExist, _ := createRandomPost(t, true)
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

// TestGetListPostsSuccessWithIsAdminTrue as find list post without exception
func TestGetListPostsSuccessWithIsAdminTrue(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()

	// Call router with argument db
	router := router.SetupRouter(db)

	// Generate random post, and get token used to create this post for check if author post is in the list
	_, strToken := createRandomPost(t, true)

	// helper.EncodedToken for generate token and get string id
	userId, _ := helper.EncodedToken(strToken)
	// Find user for get fullname
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userData, _ := userService.GetUserById(userId)

	// Create request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/posts", nil)
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
	assert.Equal(t, "List of posts", responseBody["message"])

	// Response data list posts
	var listPosts = responseBody["data"].([]interface{})
	// Response body data length not 0
	assert.NotEqual(t, 0, len(listPosts))

	// Var for count is there author who is currently login in the post list
	countAuthor := 0

	// All property not empty
	for _, list := range listPosts {
		mapList := list.(map[string]interface{})
		assert.NotEmpty(t, mapList["id"])
		assert.NotEmpty(t, mapList["author"])
		assert.NotEmpty(t, mapList["title"])
		// If author in list is same with fullname which currently login
		if mapList["author"] == userData.Fullname {
			// Increase value
			countAuthor++
		}
	}

	// var count author must be not equal 0
	assert.NotEqual(t, 0, countAuthor)
}

// TestGetListPostsSuccessWithIsAdminFalse as find list post with exception not showing post owned user is logged in
func TestGetListPostsSuccessWithIsAdminFalse(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()

	// Call router with argument db
	router := router.SetupRouter(db)

	// Generate random post, and get token used to create this post for check if author post is in the list
	_, strToken := createRandomPost(t, false)

	// helper.EncodedToken for generate token and get string id
	userId, _ := helper.EncodedToken(strToken)
	// Find user for get fullname
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userData, _ := userService.GetUserById(userId)

	// Create request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/posts", nil)
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
	assert.Equal(t, "List of posts", responseBody["message"])

	// Response data list posts
	var listPosts = responseBody["data"].([]interface{})
	// Response body data length not 0
	assert.NotEqual(t, 0, len(listPosts))

	// Var for count is there author who is currently login in the post list
	countAuthor := 0

	// All property not empty
	for _, list := range listPosts {
		mapList := list.(map[string]interface{})
		assert.NotEmpty(t, mapList["id"])
		assert.NotEmpty(t, mapList["author"])
		assert.NotEmpty(t, mapList["title"])
		// If author in list is same with fullname which currently login
		if mapList["author"] == userData.Fullname {
			// Increase value
			countAuthor++
		}
	}

	// var count author must be 0
	assert.Equal(t, 0, countAuthor)
}

// TestGetListPostsQueryWithTitle as find list post with filter by title
func TestGetListPostsQueryWithTitle(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()

	// Call router with argument db
	router := router.SetupRouter(db)

	// Generate random post, and get token used to create this post for check if author post is in the list
	newPost, strToken := createRandomPost(t, true)

	// Create request with query filter by title
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/posts?title="+newPost.Title, nil)
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
	assert.Equal(t, "List of posts", responseBody["message"])

	// Response data list posts
	var listPosts = responseBody["data"].([]interface{})
	// Response body data length not 0
	assert.NotEqual(t, 0, len(listPosts))

	// listPost the title after filter must be same title post that is search
	assert.Equal(t, listPosts[0].(map[string]interface{})["title"], newPost.Title)
}

// TestGetListPostsUnauthorized as test unauthorized
func TestGetListPostsUnauthorized(t *testing.T) {
	// Open connection to db
	db := util.SetupTestDb()

	// Call router with argument db
	router := router.SetupRouter(db)

	// Create request with query filter by title
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/posts?", nil)
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
