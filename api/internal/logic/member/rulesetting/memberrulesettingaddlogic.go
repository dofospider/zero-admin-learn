package rulesetting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingAddLogic {
	return &MemberRuleSettingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(req *types.AddMemberRuleSettingReq) (resp *types.AddMemberRuleSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
