package logic

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleRoleLogic {
	return &UpdateRoleRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleRoleLogic) UpdateRoleRole(in *sys.UpdateRoleRoleReq) (*sys.UpdateRoleRoleResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UpdateRoleRoleResp{}, nil
}
