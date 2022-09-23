package couponhistory

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryUpdateLogic {
	return &CouponHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryUpdateLogic) CouponHistoryUpdate(req *types.UpdateCouponHistoryReq) (resp *types.UpdateCouponHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
