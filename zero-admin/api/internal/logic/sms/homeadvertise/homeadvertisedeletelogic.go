package logic

import (
	"context"
	"zero-admin-learn/api/internal/common/errorx"
	"zero-admin-learn/rpc/sms/smsclient"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeAdvertiseDeleteLogic {
	return HomeAdvertiseDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(req types.DeleteHomeAdvertiseReq) (*types.DeleteHomeAdvertiseResp, error) {
	_, err := l.svcCtx.Sms.HomeAdvertiseDelete(l.ctx, &smsclient.HomeAdvertiseDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除首页广告异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除首页广告失败")
	}

	return &types.DeleteHomeAdvertiseResp{
		Code:    "000000",
		Message: "",
	}, nil
}
