package integrationconsumesetting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingListLogic {
	return &IntegrationConsumeSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingListLogic) IntegrationConsumeSettingList(req *types.ListIntegrationConsumeSettingReq) (resp *types.ListIntegrationConsumeSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
