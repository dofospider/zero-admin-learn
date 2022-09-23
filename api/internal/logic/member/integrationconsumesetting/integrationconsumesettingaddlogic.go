package integrationconsumesetting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingAddLogic {
	return &IntegrationConsumeSettingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingAddLogic) IntegrationConsumeSettingAdd(req *types.AddIntegrationConsumeSettingReq) (resp *types.AddIntegrationConsumeSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
