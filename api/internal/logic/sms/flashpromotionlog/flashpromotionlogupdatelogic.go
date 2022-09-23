package flashpromotionlog

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogUpdateLogic {
	return &FlashPromotionLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogUpdateLogic) FlashPromotionLogUpdate(req *types.UpdateFlashPromotionLogReq) (resp *types.UpdateFlashPromotionLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
