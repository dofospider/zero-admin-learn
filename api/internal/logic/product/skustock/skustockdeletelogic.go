package skustock

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockDeleteLogic {
	return &SkuStockDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockDeleteLogic) SkuStockDelete(req *types.DeleteSkuStockReq) (resp *types.DeleteSkuStockResp, err error) {
	// todo: add your logic here and delete this line

	return
}
