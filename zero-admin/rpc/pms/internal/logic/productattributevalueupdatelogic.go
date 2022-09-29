package logic

import (
	"context"
	"zero-admin-learn/rpc/model/pmsmodel"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeValueUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueUpdateLogic {
	return &ProductAttributeValueUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueUpdateLogic) ProductAttributeValueUpdate(in *pms.ProductAttributeValueUpdateReq) (*pms.ProductAttributeValueUpdateResp, error) {
	err := l.svcCtx.PmsProductAttributeValueModel.Update(pmsmodel.PmsProductAttributeValue{
		Id:                 in.Id,
		ProductId:          in.ProductId,
		ProductAttributeId: in.ProductAttributeId,
		Value:              in.Value,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeValueUpdateResp{}, nil
}
