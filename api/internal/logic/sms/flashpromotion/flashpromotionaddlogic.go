package flashpromotion

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionAddLogic {
	return &FlashPromotionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionAddLogic) FlashPromotionAdd(req *types.AddFlashPromotionReq) (resp *types.AddFlashPromotionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
