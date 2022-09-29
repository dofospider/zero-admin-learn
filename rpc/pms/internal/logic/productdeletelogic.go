package logic

import (
	"context"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductDeleteLogic) ProductDelete(in *pms.ProductDeleteReq) (*pms.ProductDeleteResp, error) {
	err := l.svcCtx.PmsProductModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductDeleteResp{}, nil
}
