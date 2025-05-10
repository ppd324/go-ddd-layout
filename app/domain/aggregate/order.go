package aggregate

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ppd324/go-ddd-layout/app/domain/entity"
)

// 聚合实体
type Order struct {
	Id         string
	UserId     int64
	Items      []*entity.OrderItem
	PayId      string
	Status     int
	TotalPrice float64
}

func NewOrder(UserId int64) *Order {
	return &Order{
		Id:     uuid.NewString(),
		UserId: UserId,
	}

}

func (o *Order) AddItem(productId string, quantity int, price float64) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity can't be zero")
	}
	newItem := &entity.OrderItem{
		ProductID: productId,
		Quantity:  quantity,
		Price:     price,
	}
	o.Items = append(o.Items, newItem)
	o.TotalPrice += float64(quantity) * price
	return nil
}
