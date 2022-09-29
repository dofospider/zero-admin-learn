package logic

import (
	"context"
	"zero-admin-learn/api/internal/common/errorx"
	"zero-admin-learn/rpc/sys/sysclient"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuDeleteLogic {
	return MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteLogic) MenuDelete(req types.DeleteMenuReq) (*types.DeleteMenuResp, error) {
	_, err := l.svcCtx.Sys.MenuDelete(l.ctx, &sysclient.MenuDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据menuId: %d,删除菜单异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除菜单失败")
	}

	return &types.DeleteMenuResp{
		Code:    "000000",
		Message: "删除菜单成功",
	}, nil
}
