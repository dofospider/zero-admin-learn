package logic

import (
	"context"

	"zero-admin-learn/rpc/oms/internal/svc"
	"zero-admin-learn/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemCheckedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemCheckedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemCheckedLogic {
	return &CartItemCheckedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemCheckedLogic) CartItemChecked(in *omsclient.CartItemCheckedReq) (*omsclient.CartItemCheckedResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.CartItemCheckedResp{}, nil
}
