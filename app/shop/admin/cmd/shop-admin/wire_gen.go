// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/server"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	dictRepo := data.NewDictRepo(dataData, logger)
	dictUsecase := biz.NewDictUsecase(dictRepo, logger)
	shopService := service.NewShopService(logger, dictUsecase)
	accountService := service.NewAccountService()
	authService := service.NewAuthService()
	httpServer := server.NewHTTPServer(confServer, logger, shopService, accountService, authService)
	grpcServer := server.NewGRPCServer(confServer, logger, shopService, accountService, authService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
