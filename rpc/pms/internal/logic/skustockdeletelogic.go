package logic

import (
	"context"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockDeleteLogic {
	return &SkuStockDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SkuStockDeleteLogic) SkuStockDelete(in *pms.SkuStockDeleteReq) (*pms.SkuStockDeleteResp, error) {
	err := l.svcCtx.PmsSkuStockModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.SkuStockDeleteResp{}, nil
}
