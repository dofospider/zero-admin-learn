package logic

import (
	"context"
	"zero-admin-learn/rpc/model/pmsmodel"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeValueAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueAddLogic {
	return &ProductAttributeValueAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueAddLogic) ProductAttributeValueAdd(in *pms.ProductAttributeValueAddReq) (*pms.ProductAttributeValueAddResp, error) {
	_, err := l.svcCtx.PmsProductAttributeValueModel.Insert(pmsmodel.PmsProductAttributeValue{
		ProductId:          in.ProductId,
		ProductAttributeId: in.ProductAttributeId,
		Value:              in.Value,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeValueAddResp{}, nil
}
