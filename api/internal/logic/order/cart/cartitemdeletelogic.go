package cart

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemDeleteLogic {
	return &CartItemDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemDeleteLogic) CartItemDelete(req *types.DeleteCartItemReq) (resp *types.DeleteCartItemResp, err error) {
	// todo: add your logic here and delete this line

	return
}
