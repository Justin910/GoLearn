package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"store/service/internal/biz"

	v1 "week04/api/store/service/v1"
)

var ProviderSet = wire.NewSet(NewStore)

type Store struct {
	v1.UnimplementedStoreServer
	log *log.Helper
	su  *biz.StoreUseCase
}

func NewStore(useCase *biz.StoreUseCase, logger log.Logger) *Store {
	return &Store{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		su:  useCase,
	}
}
