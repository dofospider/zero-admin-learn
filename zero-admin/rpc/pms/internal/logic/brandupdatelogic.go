package logic

import (
	"context"
	"zero-admin-learn/rpc/model/pmsmodel"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandUpdateLogic {
	return &BrandUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BrandUpdateLogic) BrandUpdate(in *pms.BrandUpdateReq) (*pms.BrandUpdateResp, error) {
	err := l.svcCtx.PmsBrandModel.Update(pmsmodel.PmsBrand{
		Id:                  in.Id,
		Name:                in.Name,
		FirstLetter:         in.FirstLetter,
		Sort:                in.Sort,
		FactoryStatus:       in.FactoryStatus,
		ShowStatus:          in.ShowStatus,
		ProductCount:        in.ProductCount,
		ProductCommentCount: in.ProductCommentCount,
		Logo:                in.Logo,
		BigPic:              in.BigPic,
		BrandStory:          in.BrandStory,
	})
	if err != nil {
		return nil, err
	}

	return &pms.BrandUpdateResp{}, nil
}
