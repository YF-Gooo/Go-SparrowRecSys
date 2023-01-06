package svc

import (
	"Go-SparrowRecSys/app/api/user/internal/config"
	"Go-SparrowRecSys/app/api/user/internal/middleware"
	"Go-SparrowRecSys/app/rpc/user/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	JwtAuthMiddleware rest.Middleware
	UserRpc           user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		JwtAuthMiddleware: middleware.NewJwtAuthMiddleware(c.JwtAuth.AccessSecret).Handle,
		UserRpc:           user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
