package flashpromotion

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionListLogic {
	return &FlashPromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionListLogic) FlashPromotionList(req *types.ListFlashPromotionReq) (resp *types.ListFlashPromotionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
