package operatehistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateHistoryDeleteLogic {
	return &OperateHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryDeleteLogic) OperateHistoryDelete(req *types.DeleteOperateHistoryReq) (resp *types.DeleteOperateHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
