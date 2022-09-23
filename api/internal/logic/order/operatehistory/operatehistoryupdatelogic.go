package operatehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateHistoryUpdateLogic {
	return &OperateHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryUpdateLogic) OperateHistoryUpdate(req *types.UpdateOperateHistoryReq) (resp *types.UpdateOperateHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
