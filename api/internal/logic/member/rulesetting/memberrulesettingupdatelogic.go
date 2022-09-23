package rulesetting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingUpdateLogic {
	return &MemberRuleSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingUpdateLogic) MemberRuleSettingUpdate(req *types.UpdateMemberRuleSettingReq) (resp *types.UpdateMemberRuleSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
