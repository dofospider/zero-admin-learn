package homebrand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandListLogic {
	return &HomeBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandListLogic) HomeBrandList(req *types.ListHomeBrandReq) (resp *types.ListHomeBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
