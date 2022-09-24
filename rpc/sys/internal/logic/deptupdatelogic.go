package logic

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptUpdateLogic) DeptUpdate(in *sysclient.DeptUpdateReq) (*sysclient.DeptUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.DeptUpdateResp{}, nil
}
