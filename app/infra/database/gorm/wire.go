package gorm

import "github.com/google/wire"

var ProviderSet = wire.NewSet(New, NewOrderRepository, NewUserRepository, NewProductRepository)
