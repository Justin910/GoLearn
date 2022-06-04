//go:build wireinject
// +build wireinject

package main

import (
	"store/service/internal/biz"
	"store/service/internal/conf"
	"store/service/internal/data"
	"store/service/internal/server"
	"store/service/internal/service"

	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
