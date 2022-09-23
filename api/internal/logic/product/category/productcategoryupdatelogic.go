package category

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryUpdateLogic {
	return &ProductCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate(req *types.UpdateProductCategoryReq) (resp *types.UpdateProductCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
