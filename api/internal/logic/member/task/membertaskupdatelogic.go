package task

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskUpdateLogic {
	return &MemberTaskUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskUpdateLogic) MemberTaskUpdate(req *types.UpdateMemberTaskReq) (resp *types.UpdateMemberTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
