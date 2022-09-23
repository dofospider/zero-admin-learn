package loginlog

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogListLogic {
	return &MemberLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogListLogic) MemberLoginLogList(req *types.ListMemberLoginLogReq) (resp *types.ListMemberLoginLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
