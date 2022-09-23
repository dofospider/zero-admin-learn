package level

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelListLogic {
	return &MemberLevelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelListLogic) MemberLevelList(req *types.ListMemberLevelReq) (resp *types.ListMemberLevelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
