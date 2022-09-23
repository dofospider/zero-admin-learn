package brand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductBrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductBrandUpdateLogic {
	return &ProductBrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandUpdateLogic) ProductBrandUpdate(req *types.UpdateProductBrandReq) (resp *types.UpdateProductBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
