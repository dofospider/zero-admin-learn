package logic

import (
	"context"
	"zero-admin-learn/api/internal/common/errorx"
	"zero-admin-learn/rpc/sms/smsclient"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionDeleteLogic {
	return FlashPromotionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionDeleteLogic) FlashPromotionDelete(req types.DeleteFlashPromotionReq) (*types.DeleteFlashPromotionResp, error) {
	_, err := l.svcCtx.Sms.FlashPromotionDelete(l.ctx, &smsclient.FlashPromotionDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除限时购记录异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除限时购记录失败")
	}

	return &types.DeleteFlashPromotionResp{
		Code: "000000",
		Message: "",
	}, nil
}
