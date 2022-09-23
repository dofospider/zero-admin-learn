package loginlog

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogUpdateLogic {
	return &MemberLoginLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(req *types.UpdateMemberLoginLogReq) (resp *types.UpdateMemberLoginLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
