package config

import (
	"github.com/SliverFlow/minidouyin/common/xconfig"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	DB  *xconfig.Mysql
	RDB *xconfig.Redis
}
