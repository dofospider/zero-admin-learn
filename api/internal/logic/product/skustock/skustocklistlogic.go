package skustock

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockListLogic {
	return &SkuStockListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockListLogic) SkuStockList(req *types.ListSkuStockReq) (resp *types.ListSkuStockResp, err error) {
	// todo: add your logic here and delete this line

	return
}
