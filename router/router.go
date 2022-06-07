package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/auth"
	"github.com/jabutech/simple-blog/handler"
	"github.com/jabutech/simple-blog/user"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	// Repository
	authRepository := user.NewRepository(db)

	// Service
	authService := auth.NewService(authRepository)

	// Handler
	userHandler := handler.NewAuthHandler(authService)
	pingHandler := handler.NewPingHandler()

	// Create router with gin
	router := gin.Default()
	// Use cors
	router.Use(cors.Default())

	// Router group api
	api := router.Group("/api")

	// Endpoint register
	api.POST("/register", userHandler.Register)
	// Endpoint login
	api.POST("/login", userHandler.Login)

	// Endpoint ping
	api.GET("/ping", pingHandler.Ping)

	return router
}
