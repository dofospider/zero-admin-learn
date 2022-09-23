package flashpromotionsession

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionListLogic {
	return &FlashPromotionSessionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(req *types.ListFlashPromotionSessionReq) (resp *types.ListFlashPromotionSessionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
