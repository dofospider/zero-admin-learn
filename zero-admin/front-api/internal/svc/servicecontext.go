package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-admin-learn/front-api/internal/config"
	"zero-admin-learn/rpc/oms/oms"
	"zero-admin-learn/rpc/pay/pay"
	"zero-admin-learn/rpc/pms/pms"
	"zero-admin-learn/rpc/sms/sms"
	"zero-admin-learn/rpc/sys/sys"
	"zero-admin-learn/rpc/ums/ums"
)

type ServiceContext struct {
	Config config.Config

	Sys sys.Sys
	Ums ums.Ums
	Pms pms.Pms
	Oms oms.Oms
	Sms sms.Sms
	Pay pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Sys: sys.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums: ums.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms: pms.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms: oms.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms: sms.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Pay: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
