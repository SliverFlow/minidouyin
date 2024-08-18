package svc

import (
	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/internal/config"
	"github.com/SliverFlow/minidouyin/app/user/model"
	"github.com/SliverFlow/minidouyin/common/initialize"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo model.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := initialize.InitDB(c.DB)
	if err != nil {
		logx.Errorf("[MYSQL] mysql conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}

	rdb, err := initialize.InitRDB(c.RDB)
	if err != nil {
		logx.Errorf("[REDIS] redis conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}

	userRepo := model.NewUserRepo(db, rdb)
	logx.Info("[REPO] user repo init success")

	return &ServiceContext{
		Config:   c,
		UserRepo: userRepo,
	}
}
