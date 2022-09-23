package flashpromotionlog

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogAddLogic {
	return &FlashPromotionLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogAddLogic) FlashPromotionLogAdd(req *types.AddFlashPromotionLogReq) (resp *types.AddFlashPromotionLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
