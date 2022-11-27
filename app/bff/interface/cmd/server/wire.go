//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"kratos-mono-repo/app/bff/interface/internal/biz"
	"kratos-mono-repo/app/bff/interface/internal/conf"
	"kratos-mono-repo/app/bff/interface/internal/data"
	"kratos-mono-repo/app/bff/interface/internal/server"
	"kratos-mono-repo/app/bff/interface/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
