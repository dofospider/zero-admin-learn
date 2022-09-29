package logic

import (
	"context"
	"zero-admin-learn/rpc/model/umsmodel"

	"zero-admin-learn/rpc/ums/internal/svc"
	"zero-admin-learn/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingUpdateLogic {
	return &MemberRuleSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingUpdateLogic) MemberRuleSettingUpdate(in *ums.MemberRuleSettingUpdateReq) (*ums.MemberRuleSettingUpdateResp, error) {
	err := l.svcCtx.UmsMemberRuleSettingModel.Update(umsmodel.UmsMemberRuleSetting{
		Id:                in.Id,
		ContinueSignDay:   in.ContinueSignDay,
		ContinueSignPoint: in.ContinueSignPoint,
		ConsumePerPoint:   float64(in.ConsumePerPoint),
		LowOrderAmount:    float64(in.LowOrderAmount),
		MaxPointPerOrder:  in.MaxPointPerOrder,
		Type:              in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberRuleSettingUpdateResp{}, nil
}
