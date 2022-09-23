package flashpromotionsession

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionAddLogic {
	return &FlashPromotionSessionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionAddLogic) FlashPromotionSessionAdd(req *types.AddFlashPromotionSessionReq) (resp *types.AddFlashPromotionSessionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
