package statistics

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoAddLogic {
	return &MemberStatisticsInfoAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoAddLogic) MemberStatisticsInfoAdd(req *types.AddMemberStatisticsInfoReq) (resp *types.AddMemberStatisticsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
