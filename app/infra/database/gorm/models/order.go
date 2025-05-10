package models

import "time"

type Order struct {
	Id        string    `gorm:"primary_key;column:id"`
	UserId    int64     `grom:"column:user_id;not null"`
	PayId     string    `grom:"column:pay_id;not null"` // 支付id
	Status    int       `grom:"column:status;not null"` //订单状态
	CreatedAt time.Time `grom:"column:created_at"`
	UpdatedAt time.Time `grom:"column:updated_at"`
	DeletedAt time.Time `grom:"column:deleted_at"`
}

func (Order) TableName() string {
	return "orders"
}
