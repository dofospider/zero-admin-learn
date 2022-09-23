package homebrand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandUpdateLogic {
	return &HomeBrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandUpdateLogic) HomeBrandUpdate(req *types.UpdateHomeBrandReq) (resp *types.UpdateHomeBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
