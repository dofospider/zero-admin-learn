package homeadvertise

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseUpdateLogic {
	return &HomeAdvertiseUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseUpdateLogic) HomeAdvertiseUpdate(req *types.UpdateHomeAdvertiseReq) (resp *types.UpdateHomeAdvertiseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
