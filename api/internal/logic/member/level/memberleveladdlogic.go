package level

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelAddLogic {
	return &MemberLevelAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelAddLogic) MemberLevelAdd(req *types.AddMemberLevelReq) (resp *types.AddMemberLevelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
