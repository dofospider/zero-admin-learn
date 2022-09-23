package setting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingListLogic {
	return &OrderSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingListLogic) OrderSettingList(req *types.ListOrderSettingReq) (resp *types.ListOrderSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}