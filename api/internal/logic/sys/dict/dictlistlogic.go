package dict

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictListLogic) DictList(req *types.ListDictReq) (resp *types.ListDictResp, err error) {
	// todo: add your logic here and delete this line

	return
}
