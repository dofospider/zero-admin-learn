package flashpromotionsession

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionUpdateLogic {
	return &FlashPromotionSessionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(req *types.UpdateFlashPromotionSessionReq) (resp *types.UpdateFlashPromotionSessionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
