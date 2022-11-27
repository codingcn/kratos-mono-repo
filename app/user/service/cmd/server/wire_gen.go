// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"kratos-mono-repo/app/user/service/internal/biz"
	"kratos-mono-repo/app/user/service/internal/conf"
	"kratos-mono-repo/app/user/service/internal/data"
	"kratos-mono-repo/app/user/service/internal/server"
	"kratos-mono-repo/app/user/service/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, auth *conf.Auth, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	db := data.NewGormClient(confData, logger)
	cmdable := data.NewRedisCmd(confData, logger)
	discovery := data.NewDiscovery(registry)
	orderClient := data.NewOrderServiceClient(discovery, tracerProvider)
	dataData, cleanup, err := data.NewData(confData, db, cmdable, orderClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(confServer, logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
