// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/module/payment/service/internal/biz"
	"mall-go/module/payment/service/internal/conf"
	"mall-go/module/payment/service/internal/data"
	"mall-go/module/payment/service/internal/server"
	"mall-go/module/payment/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(logger)
	if err != nil {
		return nil, nil, err
	}
	payRepo := data.NewPayRepo(dataData, logger)
	payUsecase := biz.NewPayUsecase(payRepo, logger)
	paymentService := service.NewPayService(payUsecase)
	httpServer := server.NewHTTPServer(confServer, paymentService, logger)
	grpcServer := server.NewGRPCServer(confServer, paymentService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}