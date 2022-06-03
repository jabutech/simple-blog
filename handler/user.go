package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/user"
)

type UserHandler interface {
	Register(c *gin.Context)
}

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Handler Register
func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterInput

	// Get data body from request user and passing to var input
	err := c.ShouldBindJSON(&input)
	// If error from validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map error message
		errorMessage := gin.H{"errors": errors}

		// Api Response failed with helper
		response := helper.ApiResponseFailed(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Check email is availablity
	isEmailAvailable, err := h.userService.IsEmailAvailable(input.Email)
	// If error from validation
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Server error"}
		// Api Response failed with helper
		response := helper.ApiResponseFailed(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println("DEBUG")
	fmt.Println("DEBUG")
	fmt.Println("DEBUG")
	fmt.Println("DEBUG")
	fmt.Println("DEBUG")
	fmt.Println("DEBUG")
	fmt.Println(isEmailAvailable)

	// If email is availability
	if isEmailAvailable {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Email already exist."}
		// Api Response failed with helper
		response := helper.ApiResponseFailed(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If no error, register user with service
	_, err = h.userService.Register(input)
	// If error from validation
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Server error"}
		// Api Response failed with helper
		response := helper.ApiResponseFailed(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response with helper ApiResponseSuccess
	response := helper.ApiResponseSuccess(
		http.StatusOK,
		"success",
		"You have successfully registered.",
	)

	c.JSON(http.StatusOK, response)

}
