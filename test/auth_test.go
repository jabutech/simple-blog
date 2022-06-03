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

// Test Register
func TestRegisterSuccessWithoutIsAdmin(t *testing.T) {
	// Open connection to db
	db := setupTestDb()

	// Call router with argument db
	router := setupRouter(db)

	dataBody := fmt.Sprintf(`{"fullname": "%s", "email": "%s", "password": "%s"}`, util.RandomFullname(), util.RandomEmail(), "password")

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
}
