package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"week04/app/service/store/internal/biz"
)

var _ biz.StoreRepo = (*storeRepo)(nil)

type storeRepo struct {
	data *Data
	log  *log.Helper
}

func NewStoreRepo(data *Data, logger log.Logger) biz.StoreRepo {
	return &storeRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/store")),
	}
}

func (s *storeRepo) CreateGoods(ctx context.Context, goodsInfo *biz.GoodsInfo) (int64, error) {
	return 0, nil
}
func (s *storeRepo) ListGoods(ctx context.Context, pageNo int64, pageSize int64) ([]*biz.GoodsInfo, error) {
	pageNo--
	if pageNo < 0 {
		pageNo = 0
	}

	if pageSize <= 10 {
		pageSize = 10
	}

	rows, err := s.data.db.QueryContext(ctx, "select id, num, name, detail from goods_info limit ?,?", pageNo, pageSize*pageNo)
	if err != nil {
		return nil, err
	}

	goodsInfos := make([]*biz.GoodsInfo, 0)
	var err1 error
	for rows.Next() {
		if err1 != nil {
			continue
		}

		var id int64
		var num int64
		var name string
		var detail string

		e := rows.Scan(&id, &num, &name, &detail)
		if e != nil {
			err1 = e
			continue
		}
		info := new(biz.GoodsInfo)
		info.GoodsId = id
		info.GoodsNum = num
		info.GoodsName = name
		info.GoodsDetail = detail

		goodsInfos = append(goodsInfos, info)
	}

	if err1 != nil {
		return nil, err1
	}
	return goodsInfos, nil
}
func (s *storeRepo) DecGoodsNum(ctx context.Context, goodsId int64, num int64) error {
	return nil
}
func (s *storeRepo) IncGoodsNum(ctx context.Context, goodsId int64, num int64) error {
	return nil
}
