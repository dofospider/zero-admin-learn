package loginlog

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogDeleteLogic {
	return &MemberLoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(req *types.DeleteMemberLoginLogReq) (resp *types.DeleteMemberLoginLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
