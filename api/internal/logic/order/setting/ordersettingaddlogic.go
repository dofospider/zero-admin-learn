package setting

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderSettingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingAddLogic {
	return &OrderSettingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingAddLogic) OrderSettingAdd(req *types.AddOrderSettingReq) (resp *types.AddOrderSettingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
