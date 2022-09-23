package cart

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemAddLogic {
	return &CartItemAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemAddLogic) CartItemAdd(req *types.AddCartItemReq) (resp *types.AddCartItemResp, err error) {
	// todo: add your logic here and delete this line

	return
}
