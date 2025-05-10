//go:build wireinject

package main

import (
	"github.com/google/wire"
	appservice "github.com/ppd324/go-ddd-layout/app/application/services"
	domainservice "github.com/ppd324/go-ddd-layout/app/domain/service"
	"github.com/ppd324/go-ddd-layout/app/infra/database/gorm"
	"github.com/ppd324/go-ddd-layout/app/interfaces/http"
	"github.com/ppd324/go-ddd-layout/app/interfaces/http/handlers"
	"github.com/ppd324/go-ddd-layout/app/interfaces/http/router"
	"github.com/ppd324/go-ddd-layout/conf"
)

var ProviderSet = wire.NewSet(
	conf.ProviderSet,
	gorm.ProviderSet,
	domainservice.ProviderSet,
	appservice.ProviderSet,
	handlers.ProviderSet,
	http.ProviderSet,
	router.ProviderSet,
)

func CreateApp() (*http.HttpServer, error) {
	wire.Build(ProviderSet)
	return nil, nil
}
