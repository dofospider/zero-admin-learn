package skustock

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockUpdateLogic {
	return &SkuStockUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockUpdateLogic) SkuStockUpdate(req *types.UpdateSkuStockReq) (resp *types.UpdateSkuStockResp, err error) {
	// todo: add your logic here and delete this line

	return
}
