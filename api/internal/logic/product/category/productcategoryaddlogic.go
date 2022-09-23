package category

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAddLogic {
	return &ProductCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryAddLogic) ProductCategoryAdd(req *types.AddProductCategoryReq) (resp *types.AddProductCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
