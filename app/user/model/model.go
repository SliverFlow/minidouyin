package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	// UserRepo 用户操作接口
	UserRepo interface {
		userRepo
	}

	// customUserModel 自定义用户表模型
	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserRepo 创建用户表模型
func NewUserRepo(db *gorm.DB, rdb *redis.Client) UserRepo {
	return &customUserModel{
		defaultUserModel: newUserModel(),
	}
}
