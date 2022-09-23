package job

import (
	"context"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobUpdateLogic {
	return &JobUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobUpdateLogic) JobUpdate(req *types.UpdateJobReq) (resp *types.UpdateJobResp, err error) {
	// todo: add your logic here and delete this line

	return
}
