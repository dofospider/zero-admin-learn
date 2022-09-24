package logic

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigUpdateLogic {
	return &ConfigUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigUpdateLogic) ConfigUpdate(in *sysclient.ConfigUpdateReq) (*sysclient.ConfigUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.ConfigUpdateResp{}, nil
}
