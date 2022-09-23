package couponhistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryAddLogic {
	return &CouponHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryAddLogic) CouponHistoryAdd(req *types.AddCouponHistoryReq) (resp *types.AddCouponHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
