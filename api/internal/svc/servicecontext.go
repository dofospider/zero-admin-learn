package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero-admin-learn/api/internal/config"
	"zero-admin-learn/api/internal/middleware"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
	}
}
