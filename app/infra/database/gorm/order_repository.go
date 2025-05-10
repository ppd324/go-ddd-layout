package gorm

import (
	"context"
	"github.com/ppd324/go-ddd-layout/app/domain/aggregate"
	"github.com/ppd324/go-ddd-layout/app/domain/repos"
	"github.com/ppd324/go-ddd-layout/app/infra/database/gorm/models"
	"gorm.io/gorm"
	"log"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repos.OrderRepo {
	return &OrderRepository{Db: db}
}

func (o *OrderRepository) Save(ctx context.Context, customer *aggregate.Order) error {
	// 构造DO对象
	newOrder := &models.Order{
		UserId: customer.UserId,
		PayId:  customer.PayId,
		Status: customer.Status,
	}
	err := o.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(newOrder).Error; err != nil {
			return err
		}
		orderItems := make([]*models.OrderItem, len(customer.Items))
		for i, item := range customer.Items {
			orderItems[i] = &models.OrderItem{
				OrderId:   newOrder.Id,
				ProductId: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
			}
		}
		if err := tx.Save(orderItems).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("save order failed ", err)
		return err
	}
	return nil
}

func (o *OrderRepository) QueryById(ctx context.Context, orderId string) (*aggregate.Order, error) {
	db := o.Db.WithContext(ctx)
	var order models.Order
	if err := db.Where("id = ?", orderId).First(&order).Error; err != nil {
		return nil, err
	}

	var orderItems []*models.OrderItem
	if err := db.Where("order_id = ?", orderId).Find(&orderItems).Error; err != nil {
		return nil, err
	}
	var entity aggregate.Order
	entity.Id = order.Id
	entity.Status = order.Status
	entity.UserId = order.UserId
	entity.PayId = order.PayId
	for _, orderItem := range orderItems {
		if err := entity.AddItem(orderItem.ProductId, orderItem.Quantity, orderItem.Price); err != nil {
			return nil, err
		}
	}
	return &entity, nil
}
