package logic

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *sysclient.MenuUpdateReq) (*sysclient.MenuUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.MenuUpdateResp{}, nil
}
