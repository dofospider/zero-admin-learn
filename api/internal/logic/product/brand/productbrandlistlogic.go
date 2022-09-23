package brand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductBrandListLogic {
	return &ProductBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandListLogic) ProductBrandList(req *types.ListProductBrandReq) (resp *types.ListProductBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
