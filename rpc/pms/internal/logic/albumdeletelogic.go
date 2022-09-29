package logic

import (
	"context"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumDeleteLogic {
	return &AlbumDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumDeleteLogic) AlbumDelete(in *pms.AlbumDeleteReq) (*pms.AlbumDeleteResp, error) {
	err := l.svcCtx.PmsAlbumModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.AlbumDeleteResp{}, nil
}
