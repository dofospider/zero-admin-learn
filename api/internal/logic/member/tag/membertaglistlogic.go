package tag

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagListLogic {
	return &MemberTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagListLogic) MemberTagList(req *types.ListMemberTagReq) (resp *types.ListMemberTagResp, err error) {
	// todo: add your logic here and delete this line

	return
}
