package member

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberDeleteLogic {
	return &MemberDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberDeleteLogic) MemberDelete(req *types.DeleteMemberReq) (resp *types.DeleteMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}
