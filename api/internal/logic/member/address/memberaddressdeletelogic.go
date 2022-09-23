package address

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressDeleteLogic {
	return &MemberAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressDeleteLogic) MemberAddressDelete(req *types.DeleteMemberAddressReq) (resp *types.DeleteMemberAddressResp, err error) {
	// todo: add your logic here and delete this line

	return
}
