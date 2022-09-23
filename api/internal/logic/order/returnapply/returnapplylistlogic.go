package returnapply

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyListLogic {
	return &ReturnApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyListLogic) ReturnApplyList(req *types.ListReturnApplyReq) (resp *types.ListReturnApplyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
