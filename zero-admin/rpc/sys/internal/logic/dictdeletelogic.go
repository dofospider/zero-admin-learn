package logic

import (
	"context"

	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDeleteLogic {
	return &DictDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDeleteLogic) DictDelete(in *sys.DictDeleteReq) (*sys.DictDeleteResp, error) {
	err := l.svcCtx.DictModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sys.DictDeleteResp{}, nil
}
