package integrationchangehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryListLogic {
	return &IntegrationChangeHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryListLogic) IntegrationChangeHistoryList(req *types.ListIntegrationChangeHistoryReq) (resp *types.ListIntegrationChangeHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
