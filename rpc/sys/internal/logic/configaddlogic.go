package logic

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigAddLogic {
	return &ConfigAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigAddLogic) ConfigAdd(in *sysclient.ConfigAddReq) (*sysclient.ConfigAddResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.ConfigAddResp{}, nil
}
