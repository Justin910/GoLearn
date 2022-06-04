package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

//  rpc CreateGoods(CreateGoodsReq) returns (CreateGoodsRsp);
//  rpc IncGoodsNum(IncGoodsNumReq) returns (IncGoodsNumRsp);
//  rpc ListGoods(ListGoodsReq) returns (ListGoodsRsp);
//  rpc DecGoodsNum(DecGoodsNumReq) returns (DecGoodsNumRsp);

//    int64 goods_id = 1;
//    int64 goods_num = 2;
//    string goods_name = 3;
//    string goods_detail = 4;

type GoodsInfo struct {
	GoodsId     int64
	GoodsNum    int64
	GoodsName   string
	GoodsDetail string
}

type StoreRepo interface {
	CreateGoods(ctx context.Context, goodsInfo *GoodsInfo) (int64, error)
	ListGoods(ctx context.Context, pageNo int64, pageSize int64) ([]*GoodsInfo, error)
	DecGoodsNum(ctx context.Context, goodsId int64, num int64) error
	IncGoodsNum(ctx context.Context, goodsId int64, num int64) error
}

type StoreUseCase struct {
	repo StoreRepo
	log  *log.Helper
}

func NewStoreUseCase(repo StoreRepo, logger log.Logger) *StoreUseCase {
	return &StoreUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/week04")),
	}
}

func (s *StoreUseCase) CreateGoods(ctx context.Context, info GoodsInfo) (int64, error) {
	return s.repo.CreateGoods(ctx, &info)
}

func (s *StoreUseCase) ListGoods(ctx context.Context, pageNo int64, pageSize int64) ([]*GoodsInfo, error) {
	return s.repo.ListGoods(ctx, pageNo, pageSize)
}

func (s *StoreUseCase) DecGoodsNum(ctx context.Context, goodsId int64, num int64) error {
	return s.repo.DecGoodsNum(ctx, goodsId, num)
}
func (s *StoreUseCase) IncGoodsNum(ctx context.Context, goodsId int64, num int64) error {
	return s.repo.IncGoodsNum(ctx, goodsId, num)
}
