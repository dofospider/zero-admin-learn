package growthchangehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryUpdateLogic {
	return &GrowthChangeHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryUpdateLogic) GrowthChangeHistoryUpdate(req *types.UpdateGrowthChangeHistoryReq) (resp *types.UpdateGrowthChangeHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
