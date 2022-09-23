package operatehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateHistoryListLogic {
	return &OperateHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryListLogic) OperateHistoryList(req *types.ListOperateHistoryReq) (resp *types.ListOperateHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
