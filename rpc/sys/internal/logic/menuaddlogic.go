package logic

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuAddLogic) MenuAdd(in *sysclient.MenuAddReq) (*sysclient.MenuAddResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.MenuAddResp{}, nil
}
