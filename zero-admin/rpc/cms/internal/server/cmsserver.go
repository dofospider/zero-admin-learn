// Code generated by goctl. DO NOT EDIT!
// Source: cms.proto

package server

import (
	"context"

	"zero-admin-learn/rpc/cms/cmsclient"
	"zero-admin-learn/rpc/cms/internal/logic"
	"zero-admin-learn/rpc/cms/internal/svc"
)

type CmsServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedCmsServer
}

func NewCmsServer(svcCtx *svc.ServiceContext) *CmsServer {
	return &CmsServer{
		svcCtx: svcCtx,
	}
}

func (s *CmsServer) Greet(ctx context.Context, in *cmsclient.StreamReq) (*cmsclient.StreamResp, error) {
	l := logic.NewGreetLogic(ctx, s.svcCtx)
	return l.Greet(in)
}