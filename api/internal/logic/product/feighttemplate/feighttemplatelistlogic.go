package feighttemplate

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateListLogic {
	return &FeightTemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateListLogic) FeightTemplateList(req *types.ListFeightTemplateReq) (resp *types.ListFeightTemplateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
