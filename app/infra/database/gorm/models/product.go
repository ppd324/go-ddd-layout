package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Price float64 `gorm:"column:price"`
	Name  string  `gorm:"column:name"` // 商品名称
}

func (Product) TableName() string {
	return "products"
}
