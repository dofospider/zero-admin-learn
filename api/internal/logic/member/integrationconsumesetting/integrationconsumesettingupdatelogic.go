package integrationconsumesetting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingUpdateLogic {
	return &IntegrationConsumeSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingUpdateLogic) IntegrationConsumeSettingUpdate(req *types.UpdateIntegrationConsumeSettingReq) (resp *types.UpdateIntegrationConsumeSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
