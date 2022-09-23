package homenewproduct

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeNewProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductAddLogic {
	return &HomeNewProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductAddLogic) HomeNewProductAdd(req *types.AddHomeNewProductReq) (resp *types.AddHomeNewProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
