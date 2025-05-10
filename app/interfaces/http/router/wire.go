package router

import "github.com/google/wire"

func ProvideRouters(
	orderRouter *OrderRouter,
) []Router {
	return []Router{
		orderRouter,
	}
}

var ProviderSet = wire.NewSet(
	wire.Bind(new(Router), new(*OrderRouter)),
	NewOrderRouter,
	ProvideRouters,
)
