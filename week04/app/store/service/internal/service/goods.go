package service

import (
	"context"
	"store/service/internal/biz"
	v1 "week04/api/store/service/v1"
)

func (s *Store) CreateGoods(ctx context.Context, req *v1.CreateGoodsReq) (*v1.CreateGoodsRsp, error) {
	rsp := &v1.CreateGoodsRsp{
		GoodsId: 0,
	}
	ginfo := biz.GoodsInfo{
		GoodsName:   req.GetGoodsName(),
		GoodsDetail: req.GetGoodsDetail(),
	}
	goodId, err := s.su.CreateGoods(ctx, ginfo)
	if err != nil {
		return rsp, err
	}
	rsp.GoodsId = goodId
	return rsp, nil
}
func (s *Store) IncGoodsNum(ctx context.Context, req *v1.IncGoodsNumReq) (*v1.IncGoodsNumRsp, error) {
	rsp := &v1.IncGoodsNumRsp{}
	err := s.su.IncGoodsNum(ctx, req.GetGoodsId(), req.GetGoodsNum())
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
func (s *Store) ListGoods(ctx context.Context, req *v1.ListGoodsReq) (*v1.ListGoodsRsp, error) {
	rsp := &v1.ListGoodsRsp{
		Goods: make([]*v1.ListGoodsRsp_Goods, 0),
	}
	ginfos, err := s.su.ListGoods(ctx, req.GetPageNo(), req.GetPageSize())
	if err != nil {
		return nil, err
	}

	for _, v := range ginfos {
		rsp.Goods = append(rsp.Goods, &v1.ListGoodsRsp_Goods{
			GoodsId:     v.GoodsId,
			GoodsNum:    v.GoodsNum,
			GoodsName:   v.GoodsName,
			GoodsDetail: v.GoodsDetail,
		})
	}

	return rsp, nil
}
func (s *Store) DecGoodsNum(ctx context.Context, req *v1.DecGoodsNumReq) (*v1.DecGoodsNumRsp, error) {
	rsp := &v1.DecGoodsNumRsp{}
	err := s.su.DecGoodsNum(ctx, req.GetGoodsId(), req.GetGoodsNum())
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
