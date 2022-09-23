package flashpromotion

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionUpdateLogic {
	return &FlashPromotionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionUpdateLogic) FlashPromotionUpdate(req *types.UpdateFlashPromotionReq) (resp *types.UpdateFlashPromotionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
