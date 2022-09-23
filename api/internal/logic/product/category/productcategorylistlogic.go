package category

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryListLogic {
	return &ProductCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryListLogic) ProductCategoryList(req *types.ListProductCategoryReq) (resp *types.ListProductCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
