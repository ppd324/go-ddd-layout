package handlers

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewOrderHandler)
