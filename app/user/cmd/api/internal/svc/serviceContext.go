package svc

import (
	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/config"
	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/pb"
	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc pb.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: newUserRpcClient(c.UserRpc),
	}
}

// newUserRpcClient
func newUserRpcClient(c zrpc.RpcClientConf) pb.UserClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] user rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS] user rpc client conn success \n")
	return user.NewUser(client)
}
