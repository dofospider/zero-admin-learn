package config

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigUpdateLogic {
	return &ConfigUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigUpdateLogic) ConfigUpdate(req *types.UpdateConfigReq) (resp *types.UpdateConfigResp, err error) {
	// todo: add your logic here and delete this line

	return
}
