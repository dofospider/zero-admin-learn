package logic

import (
	"context"
	"zero-admin-learn/rpc/model/pmsmodel"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFullReductionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionAddLogic {
	return &ProductFullReductionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionAddLogic) ProductFullReductionAdd(in *pms.ProductFullReductionAddReq) (*pms.ProductFullReductionAddResp, error) {
	_, err := l.svcCtx.PmsProductFullReductionModel.Insert(pmsmodel.PmsProductFullReduction{
		ProductId:   in.ProductId,
		FullPrice:   float64(in.FullPrice),
		ReducePrice: float64(in.ReducePrice),
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductFullReductionAddResp{}, nil
}
