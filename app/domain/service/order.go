package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ppd324/go-ddd-layout/app/domain/aggregate"
	"github.com/ppd324/go-ddd-layout/app/domain/repos"
	"gorm.io/gorm"
)

type IOrderService interface {
}

type OrderService struct {
	orderRepo   repos.OrderRepo
	userRepo    repos.UserRepo
	productRepo repos.ProductRepo
}

func NewOrderService(orderRepo repos.OrderRepo, userRepo repos.UserRepo, productRepo repos.ProductRepo) *OrderService {
	return &OrderService{orderRepo, userRepo, productRepo}
}

func (o *OrderService) ValidateOrder(ctx context.Context, order *aggregate.Order) error {
	//查询是否存在对应用户以及商品
	user, err := o.userRepo.GetById(ctx, int(order.UserId))
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user not found")
	}
	if user.Id <= 0 {
		return fmt.Errorf("user not found")
	}
	productIds := make([]string, len(order.Items))
	for i, item := range order.Items {
		productIds[i] = item.ProductID
	}
	if exist, notExistId, _ := o.productRepo.Exist(ctx, productIds); !exist {
		return fmt.Errorf("product %v not found", notExistId)
	}
	return nil
}

func (o *OrderService) CreateOrder(ctx context.Context, order *aggregate.Order) error {
	return o.orderRepo.Save(ctx, order)
}
