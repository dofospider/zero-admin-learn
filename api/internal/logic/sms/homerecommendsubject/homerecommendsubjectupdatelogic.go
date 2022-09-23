package homerecommendsubject

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectUpdateLogic {
	return &HomeRecommendSubjectUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectUpdateLogic) HomeRecommendSubjectUpdate(req *types.UpdateHomeRecommendSubjectReq) (resp *types.UpdateHomeRecommendSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
