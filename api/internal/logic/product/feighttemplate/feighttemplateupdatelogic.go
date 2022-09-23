package feighttemplate

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateUpdateLogic {
	return &FeightTemplateUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateUpdateLogic) FeightTemplateUpdate(req *types.UpdateFeightTemplateReq) (resp *types.UpdateFeightTemplateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
