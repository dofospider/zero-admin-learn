package homerecommendsubject

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectDeleteLogic {
	return &HomeRecommendSubjectDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectDeleteLogic) HomeRecommendSubjectDelete(req *types.DeleteHomeRecommendSubjectReq) (resp *types.DeleteHomeRecommendSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
