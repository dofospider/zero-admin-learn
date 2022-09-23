package returnapply

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyAddLogic {
	return &ReturnApplyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyAddLogic) ReturnApplyAdd(req *types.AddReturnApplyReq) (resp *types.AddReturnApplyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
