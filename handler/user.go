package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/user"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

// GetUsers find all users data
func (h *userHandler) GetUsers(c *gin.Context) {
	// Get query `fullname` and `email`
	fullname := c.Query("fullname")
	email := c.Query("email")

	// Get user
	users, err := h.service.GetUsers(fullname, email)
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Server error"}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Error to get users",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response with helper ApiResponseWithData
	response := helper.ApiResponseWithData(
		http.StatusOK,
		"success",
		"List of users",
		user.FormatUsers(users), // use formatter
	)

	c.JSON(http.StatusOK, response)
}

// GetUsers find all users data
func (h *userHandler) GetUser(c *gin.Context) {
	var userId user.GetIdUserInput

	// Get uri `id`
	err := c.ShouldBindUri(&userId)
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": err}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Failed get uri id",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get user
	userData, err := h.service.GetUserById(userId.Id)
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Server error"}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Error to get user",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if userData.ID == "" {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "User not found"}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Error to get user detail",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response with helper ApiResponseWithData
	response := helper.ApiResponseWithData(
		http.StatusOK,
		"success",
		"User detail",
		user.FormatUser(userData), // use formatter
	)

	c.JSON(http.StatusOK, response)
}
