package compayaddress

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompayAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompayAddressDeleteLogic {
	return &CompayAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressDeleteLogic) CompayAddressDelete(req *types.DeleteCompayAddressReq) (resp *types.DeleteCompayAddressResp, err error) {
	// todo: add your logic here and delete this line

	return
}
