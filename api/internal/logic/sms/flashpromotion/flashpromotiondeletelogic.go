package flashpromotion

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionDeleteLogic {
	return &FlashPromotionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionDeleteLogic) FlashPromotionDelete(req *types.DeleteFlashPromotionReq) (resp *types.DeleteFlashPromotionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
