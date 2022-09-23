package statistics

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoUpdateLogic {
	return &MemberStatisticsInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoUpdateLogic) MemberStatisticsInfoUpdate(req *types.UpdateMemberStatisticsInfoReq) (resp *types.UpdateMemberStatisticsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
