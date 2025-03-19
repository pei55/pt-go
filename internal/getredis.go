/**
*@Author: pei5
*@Date: 2024/1/16 10:16 AM
*@File: /getredis.go
*@Version:
*@Description:
**/
package internal

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", MyAppConf.RedisInfo.Host, MyAppConf.RedisInfo.Port),
		Password: MyAppConf.RedisInfo.Password,
		DB:       MyAppConf.RedisInfo.Db,
	})

	return rdb
}
