package feighttemplate

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateAddLogic {
	return &FeightTemplateAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateAddLogic) FeightTemplateAdd(req *types.AddFeightTemplateReq) (resp *types.AddFeightTemplateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
