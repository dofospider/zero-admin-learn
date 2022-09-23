package job

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobAddLogic {
	return &JobAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobAddLogic) JobAdd(req *types.AddJobReq) (resp *types.AddJobResp, err error) {
	// todo: add your logic here and delete this line

	return
}
