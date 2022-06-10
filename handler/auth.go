package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/auth"
	"github.com/jabutech/simple-blog/helper"
)

type authHandler struct {
	authService auth.Service
}

func NewAuthHandler(authService auth.Service) *authHandler {
	return &authHandler{authService}
}

// Handler Register
func (h *authHandler) Register(c *gin.Context) {
	var input auth.RegisterInput

	// Get data body from request user and passing to var input
	err := c.ShouldBindJSON(&input)
	// If error from validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map error message
		errorMessage := gin.H{"errors": errors}

		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Check email is availablity
	isEmailAvailable, err := h.authService.IsEmailAvailable(input.Email)
	// If error from validation
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Server error"}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If email is availability
	if isEmailAvailable {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Email already exist."}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If no error, register user with service
	_, err = h.authService.Register(input)
	// If error from validation
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Server error"}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Registered failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response with helper ApiResponseWithoutData
	response := helper.ApiResponseWithoutData(
		http.StatusOK,
		"success",
		"You have successfully registered.",
	)

	c.JSON(http.StatusOK, response)
}

func (h *authHandler) Login(c *gin.Context) {
	var input auth.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map error message
		errorMessage := gin.H{"errors": errors}

		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Login failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Process login
	loggedinUSer, err := h.authService.Login(input)
	if err != nil {
		// Create new map error message
		errorMessage := gin.H{"errors": err.Error()}

		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusUnprocessableEntity,
			"error",
			"Login failed.",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Generate token
	token, err := h.authService.GenerateToken(loggedinUSer)
	if err != nil {
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Login failed.",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	fieldToken := gin.H{"token": token}
	// Create format response
	response := helper.ApiResponseWithData(
		http.StatusOK,
		"success",
		"You have Login.",
		fieldToken,
	)

	c.JSON(http.StatusOK, response)

}
