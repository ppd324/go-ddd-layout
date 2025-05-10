package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ppd324/go-ddd-layout/app/application/dto"
	"github.com/ppd324/go-ddd-layout/app/application/services"
	"github.com/ppd324/go-ddd-layout/app/interfaces/http/response"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req dto.CreateOrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.ParameterError(ctx, err.Error(), nil)
		return
	}
	err := h.orderService.CreateOrder(ctx, &req)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Ok(ctx, nil)
}
