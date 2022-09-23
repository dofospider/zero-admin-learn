package skustock

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockAddLogic {
	return &SkuStockAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockAddLogic) SkuStockAdd(req *types.AddSkuStockReq) (resp *types.AddSkuStockResp, err error) {
	// todo: add your logic here and delete this line

	return
}
