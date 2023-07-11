package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler interface {
	Home(c *gin.Context)
	PingPong(c *gin.Context)
}

type baseHandler struct{}

func NewBaseHandler() BaseHandler {
	return &baseHandler{}
}

func (h *baseHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Hello World",
		"data": gin.H{
			"name":     "Muhamad Rizal Arfiyan",
			"username": "rizalarfiyan",
			"email":    "rizal.arfiyan.23@gmail.com",
			"github":   "https://github.com/rizalarfiyan/",
			"linkedin": "https://www.linkedin.com/in/rizalarfiyan/",
		},
	})
}
func (h *baseHandler) PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "pong",
		"data":    nil,
	})
}
