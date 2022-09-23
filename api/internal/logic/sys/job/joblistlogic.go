package job

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobListLogic {
	return &JobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobListLogic) JobList(req *types.ListJobReq) (resp *types.ListJobResp, err error) {
	// todo: add your logic here and delete this line

	return
}
