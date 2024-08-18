package initialize

import (
	"context"
	"fmt"
	"github.com/SliverFlow/minidouyin/common/xconfig"
	"github.com/redis/go-redis/v9"
)

func InitRDB(c *xconfig.Redis) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		Password: c.Password, // 没有密码，默认值
		DB:       c.Db,       // 默认DB 0
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
