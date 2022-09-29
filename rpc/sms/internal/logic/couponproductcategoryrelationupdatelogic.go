package logic

import (
	"context"
	"zero-admin-learn/rpc/model/smsmodel"

	"zero-admin-learn/rpc/sms/internal/svc"
	"zero-admin-learn/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductCategoryRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationUpdateLogic {
	return &CouponProductCategoryRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationUpdateLogic) CouponProductCategoryRelationUpdate(in *sms.CouponProductCategoryRelationUpdateReq) (*sms.CouponProductCategoryRelationUpdateResp, error) {
	err := l.svcCtx.SmsCouponProductCategoryRelationModel.Update(smsmodel.SmsCouponProductCategoryRelation{
		Id:                  in.Id,
		CouponId:            in.CouponId,
		ProductCategoryId:   in.ProductCategoryId,
		ProductCategoryName: in.ProductCategoryName,
		ParentCategoryName:  in.ParentCategoryName,
	})
	if err != nil {
		return nil, err
	}

	return &sms.CouponProductCategoryRelationUpdateResp{}, nil
}
