package router

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/auth"
	"github.com/jabutech/simple-blog/handler"
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/user"
	"github.com/jabutech/simple-blog/util"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	// Router
	router := NewRouter(db)

	return router
}

func NewRouter(db *gorm.DB) *gin.Engine {
	// Repository
	userRepository := user.NewRepository(db)

	// Service
	authService := auth.NewService(userRepository)
	userService := user.NewService(userRepository)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	pingHandler := handler.NewPingHandler()

	// Create router with gin
	router := gin.Default()
	// Use cors
	router.Use(cors.Default())

	// Router group api
	api := router.Group("/api")

	// Endpoint register
	api.POST("/register", authHandler.Register)
	// Endpoint login
	api.POST("/login", authHandler.Login)

	// Group users
	users := api.Group("/users")
	// Endpoint get all users
	users.GET("", authMiddleware(userService), userHandler.GetUsers)
	// Endpoint get all user by id
	users.GET("/:id", authMiddleware(userService), userHandler.GetUser)

	// Endpoint ping
	api.GET("/ping", pingHandler.Ping)

	return router
}

// Function for handle middleware
func authMiddleware(userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get header with name `Authorization`
		authHeader := c.GetHeader("Authorization")

		// If inside authHeader doesn't have `Bearer`
		if !strings.Contains(authHeader, "Bearer") {
			// Create format response with helper
			response := helper.ApiResponseWithoutData(http.StatusUnauthorized, "error", "Unauthorized")
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// If there is, create new variable with empty string value
		encodedToken := ""
		// Split authHeader with white space
		arrayToken := strings.Split(authHeader, " ")
		// If length arrayToken is same the 2
		if len(arrayToken) == 2 {
			// Get arrayToken with index 1 / only token jwt
			encodedToken = arrayToken[1]
		}

		// Load config
		config, err := util.LoadConfig("../") // "." as location file app.env in root folder
		helper.FatalError("error load config: ", err)

		// Validation token
		token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("Invalid token")
			}

			return []byte(config.SecretKey), nil
		})

		// If error
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			// Create format response with helper
			response := helper.ApiResponseWithData(http.StatusUnauthorized, "error", "Unauthorized", errorMessage)
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Get payload token
		claim, ok := token.Claims.(jwt.MapClaims)
		// If not `ok` and token invalid
		if !ok || !token.Valid {
			// Create format response with helper
			response := helper.ApiResponseWithoutData(http.StatusUnauthorized, "error", "Unauthorized")
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Get payload `id` and convert to type `float64` and type `int`
		userId := fmt.Sprint(claim["user_id"])

		// Find user on db with service
		user, err := userService.GetUserById(userId)
		// If error
		if err != nil {
			// Create format response with helper
			response := helper.ApiResponseWithoutData(http.StatusUnauthorized, "error", "Unauthorized")
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Set user to context with name `currentUser`
		c.Set("currentUser", user)
	}
}
