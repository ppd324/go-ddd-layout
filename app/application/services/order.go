package services

import (
	"context"
	"github.com/ppd324/go-ddd-layout/app/application/dto"
	"github.com/ppd324/go-ddd-layout/app/domain/aggregate"
	domain "github.com/ppd324/go-ddd-layout/app/domain/service"
)

type OrderService struct {
	domainService *domain.OrderService
}

func NewOrderService(server *domain.OrderService) *OrderService {
	return &OrderService{domainService: server}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error {
	newOrder := aggregate.NewOrder(req.UserId)
	if err := o.domainService.ValidateOrder(ctx, newOrder); err != nil {
		return err
	}
	return o.domainService.CreateOrder(ctx, newOrder)
}
