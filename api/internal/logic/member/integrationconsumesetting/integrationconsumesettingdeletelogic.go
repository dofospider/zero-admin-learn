package integrationconsumesetting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingDeleteLogic {
	return &IntegrationConsumeSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(req *types.DeleteIntegrationConsumeSettingReq) (resp *types.DeleteIntegrationConsumeSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
