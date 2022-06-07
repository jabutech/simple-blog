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

	// TODO: authorization
}
