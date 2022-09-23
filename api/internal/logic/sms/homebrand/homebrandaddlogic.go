package homebrand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandAddLogic {
	return &HomeBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandAddLogic) HomeBrandAdd(req *types.AddHomeBrandReq) (resp *types.AddHomeBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
