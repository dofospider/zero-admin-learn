package statistics

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoListLogic {
	return &MemberStatisticsInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoListLogic) MemberStatisticsInfoList(req *types.ListMemberStatisticsInfoReq) (resp *types.ListMemberStatisticsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
