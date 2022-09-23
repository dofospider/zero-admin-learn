package homerecommendsubject

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectListLogic {
	return &HomeRecommendSubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectListLogic) HomeRecommendSubjectList(req *types.ListHomeRecommendSubjectReq) (resp *types.ListHomeRecommendSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
