package handler

import "github.com/gin-gonic/gin"

type PingHandler interface {
	Ping(c *gin.Context)
}

type pingHandler struct {
}

// Instance
func NewPingHandler() *pingHandler {
	return &pingHandler{}
}

// Function ping
func (h *pingHandler) Ping(c *gin.Context) {
	response := gin.H{"status": "pong"}
	c.JSON(200, response)
}
