package homerecommendproduct

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductAddLogic {
	return &HomeRecommendProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(req *types.AddHomeRecommendProductReq) (resp *types.AddHomeRecommendProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
