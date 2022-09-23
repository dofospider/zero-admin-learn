package comment

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommentDeleteLogic {
	return &ProductCommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentDeleteLogic) ProductCommentDelete(req *types.DeleteProductCommentReq) (resp *types.DeleteProductCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
