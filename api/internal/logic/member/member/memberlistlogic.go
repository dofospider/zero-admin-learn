package member

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberListLogic {
	return &MemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberListLogic) MemberList(req *types.ListMemberReq) (resp *types.ListMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}
