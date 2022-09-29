package logic

import (
	"context"
	"zero-admin-learn/rpc/model/umsmodel"

	"zero-admin-learn/rpc/ums/internal/svc"
	"zero-admin-learn/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingAddLogic {
	return &MemberRuleSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(in *ums.MemberRuleSettingAddReq) (*ums.MemberRuleSettingAddResp, error) {
	_, err := l.svcCtx.UmsMemberRuleSettingModel.Insert(umsmodel.UmsMemberRuleSetting{
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

	return &ums.MemberRuleSettingAddResp{}, nil
}
