package growthchangehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryDeleteLogic {
	return &GrowthChangeHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryDeleteLogic) GrowthChangeHistoryDelete(req *types.DeleteGrowthChangeHistoryReq) (resp *types.DeleteGrowthChangeHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
