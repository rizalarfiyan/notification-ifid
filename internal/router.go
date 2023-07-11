package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/rizalarfiyan/notification-ifid/internal/handler"
	"github.com/rizalarfiyan/notification-ifid/internal/services"
)

type Router interface {
	All()
}

type router struct {
	route *gin.Engine
}

func NewRouter(route *gin.Engine) Router {
	return &router{route}
}

func (r *router) All() {
	r.base()
	r.notification()
}

func (r *router) base() {
	base := handler.NewBaseHandler()
	r.route.GET("/", base.Home)
	r.route.GET("/ping", base.PingPong)
}

func (r *router) notification() {
	service := services.NewNotificationService()
	notification := handler.NewNotificationHandler(service)
	r.route.POST("/notification", notification.SendNotification)
}
