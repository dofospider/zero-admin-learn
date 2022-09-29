package logic

import (
	"context"
	"zero-admin-learn/rpc/model/pmsmodel"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeightTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateUpdateLogic {
	return &FeightTemplateUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeightTemplateUpdateLogic) FeightTemplateUpdate(in *pms.FeightTemplateUpdateReq) (*pms.FeightTemplateUpdateResp, error) {
	err := l.svcCtx.PmsFeightTemplateModel.Update(pmsmodel.PmsFeightTemplate{
		Id:             in.Id,
		Name:           in.Name,
		ChargeType:     in.ChargeType,
		FirstWeight:    float64(in.FirstWeight),
		FirstFee:       float64(in.FirstFee),
		ContinueWeight: float64(in.ContinueWeight),
		ContinmeFee:    float64(in.ContinmeFee),
		Dest:           in.Dest,
	})
	if err != nil {
		return nil, err
	}

	return &pms.FeightTemplateUpdateResp{}, nil
}
