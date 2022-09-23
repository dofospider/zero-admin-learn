package memberprice

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPriceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceDeleteLogic {
	return &MemberPriceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceDeleteLogic) MemberPriceDelete(req *types.DeleteMemberPriceReq) (resp *types.DeleteMemberPriceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
