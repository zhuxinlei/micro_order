package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct{
		DataSource string
	}


	Auth      struct {
		AccessSecret string
		AccessExpire int64
	}
	BookRpc zrpc.RpcClientConf
}
