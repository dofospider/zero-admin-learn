package logic

import (
	"context"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryAttributeRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationDeleteLogic {
	return &ProductCategoryAttributeRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationDeleteLogic) ProductCategoryAttributeRelationDelete(in *pms.ProductCategoryAttributeRelationDeleteReq) (*pms.ProductCategoryAttributeRelationDeleteResp, error) {
	err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductCategoryAttributeRelationDeleteResp{}, nil
}
