package returnreason

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnResonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnResonListLogic {
	return &ReturnResonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonListLogic) ReturnResonList(req *types.ListReturnResonReq) (resp *types.ListReturnResonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
