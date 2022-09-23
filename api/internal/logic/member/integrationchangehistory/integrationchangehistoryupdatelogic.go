package integrationchangehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryUpdateLogic {
	return &IntegrationChangeHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryUpdateLogic) IntegrationChangeHistoryUpdate(req *types.UpdateIntegrationChangeHistoryReq) (resp *types.UpdateIntegrationChangeHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
