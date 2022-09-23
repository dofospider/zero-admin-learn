package brand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductBrandAddLogic {
	return &ProductBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandAddLogic) ProductBrandAdd(req *types.AddProductBrandReq) (resp *types.AddProductBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
