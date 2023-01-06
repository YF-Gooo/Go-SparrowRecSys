package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Domain  string
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRpcConf zrpc.RpcClientConf
}
