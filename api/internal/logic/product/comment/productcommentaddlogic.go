package comment

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommentAddLogic {
	return &ProductCommentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentAddLogic) ProductCommentAdd(req *types.AddProductCommentReq) (resp *types.AddProductCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
