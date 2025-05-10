package repos

import (
	"context"
	"github.com/ppd324/go-ddd-layout/app/domain/aggregate"
)

type OrderRepo interface {
	// 创建订单
	Save(ctx context.Context, customer *aggregate.Order) error
	// 查询订单
	QueryById(ctx context.Context, orderId string) (*aggregate.Order, error)
}
