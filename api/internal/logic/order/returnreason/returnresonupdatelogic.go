package returnreason

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnResonUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnResonUpdateLogic {
	return &ReturnResonUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonUpdateLogic) ReturnResonUpdate(req *types.UpdateReturnResonReq) (resp *types.UpdateReturnResonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
