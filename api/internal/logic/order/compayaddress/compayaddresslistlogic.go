package compayaddress

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompayAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompayAddressListLogic {
	return &CompayAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressListLogic) CompayAddressList(req *types.ListCompayAddressReq) (resp *types.ListCompayAddressResp, err error) {
	// todo: add your logic here and delete this line

	return
}
