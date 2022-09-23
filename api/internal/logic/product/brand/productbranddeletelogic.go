package brand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductBrandDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductBrandDeleteLogic {
	return &ProductBrandDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandDeleteLogic) ProductBrandDelete(req *types.DeleteProductBrandReq) (resp *types.DeleteProductBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
