package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"greet/services/user/userapi/internal/config"
	"greet/services/user/userrpc/userrpcclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userrpcclient.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpcclient.NewUserrpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
