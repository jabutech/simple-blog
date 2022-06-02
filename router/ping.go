package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jabutech/simple-blog/handler"
)

func NewPingRouter(pingHandler handler.PingHandler) *gin.Engine {
	// Create router with gin
	router := gin.Default()
	// Router group api
	api := router.Group("/api")

	// Endpoint ping
	api.GET("/ping", pingHandler.Ping)

	return router
}
