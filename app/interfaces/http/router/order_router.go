package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ppd324/go-ddd-layout/app/interfaces/http/handlers"
	"github.com/ppd324/go-ddd-layout/app/interfaces/http/middlewares"
)

type OrderRouter struct {
	handler *handlers.OrderHandler
}

func NewOrderRouter(handler *handlers.OrderHandler) *OrderRouter {
	return &OrderRouter{
		handler: handler,
	}

}

func (r *OrderRouter) Name() string {
	return "OrderRouter"
}

func (r *OrderRouter) Mount(gin *gin.Engine) {
	orderRouter := gin.Group("/orders", middlewares.CorsMiddleware())
	{
		orderRouter.POST("", r.handler.CreateOrder)
	}

}
