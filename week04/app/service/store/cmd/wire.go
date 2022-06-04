//go:build wireinject
// +build wireinject

package main

import (
	"week04/app/service/store/internal/biz"
	"week04/app/service/store/internal/conf"
	"week04/app/service/store/internal/data"
	"week04/app/service/store/internal/server"
	"week04/app/service/store/internal/service"

	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
