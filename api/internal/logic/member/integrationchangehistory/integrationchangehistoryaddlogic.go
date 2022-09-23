package integrationchangehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryAddLogic {
	return &IntegrationChangeHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryAddLogic) IntegrationChangeHistoryAdd(req *types.AddIntegrationChangeHistoryReq) (resp *types.AddIntegrationChangeHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
