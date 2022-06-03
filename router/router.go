package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/handler"
)

func NewRouter(pingHandler handler.PingHandler, userHandler handler.UserHandler) *gin.Engine {
	// Create router with gin
	router := gin.Default()
	// Router group api
	api := router.Group("/api")

	// Router group users
	users := api.Group("/users")
	// Endpoint users
	users.POST("", userHandler.Register)

	// Endpoint ping
	api.GET("/ping", pingHandler.Ping)

	return router
}
