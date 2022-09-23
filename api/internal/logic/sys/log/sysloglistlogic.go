package log

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogListLogic) SysLogList(req *types.ListSysLogReq) (resp *types.ListSysLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
