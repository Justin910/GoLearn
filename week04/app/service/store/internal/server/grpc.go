package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "week04/api/store/service/v1"
	"week04/app/service/store/internal/conf"
	"week04/app/service/store/internal/service"
)

func NewGRPCServer(c *conf.Server, logger log.Logger, s *service.Store) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(*c.Grpc.Timeout))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterStoreServer(srv, s)
	return srv
}
