package models

import (
	"github.com/ppd324/go-ddd-layout/app/domain/entity"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderId   string  `gorm:"column:order_id;not null"`
	ProductId string  `gorm:"column:product_id;type:int(11);not null"`
	Quantity  int     `gorm:"column:quantity;type:int(11);not null"`
	Price     float64 `gorm:"column:price"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

func (o *OrderItem) ToEntity() *entity.OrderItem {
	return &entity.OrderItem{
		ProductID: o.ProductId,
		Quantity:  o.Quantity,
		Price:     o.Price,
	}
}
