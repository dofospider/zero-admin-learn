package growthchangehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryAddLogic {
	return &GrowthChangeHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryAddLogic) GrowthChangeHistoryAdd(req *types.AddGrowthChangeHistoryReq) (resp *types.AddGrowthChangeHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
