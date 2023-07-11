package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizalarfiyan/notification-ifid/internal/request"
	"github.com/rizalarfiyan/notification-ifid/internal/services"
)

type NotificationHandler interface {
	SendNotification(c *gin.Context)
}

type notificationHandler struct {
	service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) NotificationHandler {
	return &notificationHandler{service}
}

func (h *notificationHandler) SendNotification(c *gin.Context) {
	var params request.NotificationRequest
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error: " + err.Error(),
			"data":    nil,
		})
		return
	}

	err = h.service.SendNotification(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success send notification!",
		"data":    nil,
	})
}
