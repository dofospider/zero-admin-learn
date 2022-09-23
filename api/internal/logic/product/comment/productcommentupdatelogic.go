package comment

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommentUpdateLogic {
	return &ProductCommentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentUpdateLogic) ProductCommentUpdate(req *types.UpdateProductCommentReq) (resp *types.UpdateProductCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
