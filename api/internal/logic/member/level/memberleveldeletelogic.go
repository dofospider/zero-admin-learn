package level

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelDeleteLogic {
	return &MemberLevelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelDeleteLogic) MemberLevelDelete(req *types.DeleteMemberLevelReq) (resp *types.DeleteMemberLevelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
