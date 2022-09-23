package setting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingDeleteLogic {
	return &OrderSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingDeleteLogic) OrderSettingDelete(req *types.DeleteOrderSettingReq) (resp *types.DeleteOrderSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
