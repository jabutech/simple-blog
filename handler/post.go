package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/post"
	"github.com/jabutech/simple-blog/user"
)

type PostHandler interface {
	Create(c *gin.Context)
	GetPosts(c *gin.Context)
}

type postHandler struct {
	postService post.Service
}

func NewPostHandler(postService post.Service) *postHandler {
	return &postHandler{postService}
}

// Create for create new post
func (h *postHandler) Create(c *gin.Context) {
	var input post.CreatePostInput

	// Get data current user is logged in from context
	currentUser := c.MustGet("currentUser").(user.User)
	// Passing id current user into var input.UserId
	input.UserId = currentUser.ID

	// Get data body from request
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
			"Failed create post",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create new post
	newPost, err := h.postService.Create(input)
	// If error from validation
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": err.Error()}
		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Failed create post",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Format response
	formatPost := post.FormatPostCreateOrUpdate(newPost)

	// Create format response with helper ApiResponseWithoutData
	response := helper.ApiResponseWithData(
		http.StatusOK,
		"success",
		"Post has been created",
		formatPost,
	)

	c.JSON(http.StatusOK, response)
}

func (h *postHandler) Update(c *gin.Context) {
	var updateInput post.UpdatePostInput

	// get request title from json
	err := c.ShouldBindJSON(&updateInput)
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map error message
		errorMessage := gin.H{"errors": errors}

		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Updated failed",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get id from uri
	err = c.ShouldBindUri(&updateInput)
	if err != nil {
		// Create new map error message
		errorMessage := gin.H{"errors": err}

		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Updated failed",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get data current user is logged in from context
	currentUser := c.MustGet("currentUser").(user.User)
	// Passing id current user into var input.UserId
	updateInput.UserId = currentUser.ID

	// Update
	updatedPost, err := h.postService.Update(updateInput)
	if err != nil {
		// Create new map error message
		errorMessage := gin.H{"errors": err.Error()}

		// Api Response failed with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Updated failed",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Format response
	formatPost := post.FormatPostCreateOrUpdate(updatedPost)

	// Create format response with helper ApiResponseWithoutData
	response := helper.ApiResponseWithData(
		http.StatusOK,
		"success",
		"Post has been updated successfully",
		formatPost,
	)

	c.JSON(http.StatusOK, response)
}

func (h *postHandler) GetPosts(c *gin.Context) {
	// Get query `title`
	title := c.Query("title")

	// Get data current user is logged in from context
	currentUser := c.MustGet("currentUser").(user.User)

	// Get all post
	posts, err := h.postService.GetPosts(title, currentUser)
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "server error"}
		// Create response with helper
		response := helper.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Error to get posts",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If no error, create format with helper
	formatter := post.FormatPosts(posts)
	// Create response with helper
	response := helper.ApiResponseWithData(
		http.StatusOK,
		"success",
		"List of posts",
		formatter,
	)

	c.JSON(http.StatusOK, response)
	return

}
