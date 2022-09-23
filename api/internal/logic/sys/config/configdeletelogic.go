package config

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigDeleteLogic {
	return &ConfigDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigDeleteLogic) ConfigDelete(req *types.DeleteConfigReq) (resp *types.DeleteConfigResp, err error) {
	// todo: add your logic here and delete this line

	return
}
