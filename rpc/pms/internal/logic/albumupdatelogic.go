package logic

import (
	"context"
	"zero-admin-learn/rpc/model/pmsmodel"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumUpdateLogic {
	return &AlbumUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumUpdateLogic) AlbumUpdate(in *pms.AlbumUpdateReq) (*pms.AlbumUpdateResp, error) {
	err := l.svcCtx.PmsAlbumModel.Update(pmsmodel.PmsAlbum{
		Id:          in.Id,
		Name:        in.Name,
		CoverPic:    in.CoverPic,
		PicCount:    in.PicCount,
		Sort:        in.Sort,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pms.AlbumUpdateResp{}, nil
}
