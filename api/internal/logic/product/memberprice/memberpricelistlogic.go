package memberprice

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPriceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceListLogic {
	return &MemberPriceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceListLogic) MemberPriceList(req *types.ListMemberPriceReq) (resp *types.ListMemberPriceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
