package logic

import (
	"context"
	"zero-admin-learn/api/internal/common/errorx"
	"zero-admin-learn/rpc/ums/umsclient"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoDeleteLogic {
	return MemberStatisticsInfoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoDeleteLogic) MemberStatisticsInfoDelete(req types.DeleteMemberStatisticsInfoReq) (*types.DeleteMemberStatisticsInfoResp, error) {
	_, err := l.svcCtx.Ums.MemberStatisticsInfoDelete(l.ctx, &umsclient.MemberStatisticsInfoDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员统计信息异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员统计信息失败")
	}
	return &types.DeleteMemberStatisticsInfoResp{
		Code:    "000000",
		Message: "",
	}, nil
}
