package homebrand

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBrandDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandDeleteLogic {
	return &HomeBrandDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandDeleteLogic) HomeBrandDelete(req *types.DeleteHomeBrandReq) (resp *types.DeleteHomeBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
