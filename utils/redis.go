package utils

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var RedisCtx context.Context

func InitRedis(addr string, password string, db int) (err error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	RedisCtx = context.Background()
	pong, err := Redis.Ping(RedisCtx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接Redis成功:", pong) // 输出: 连接Redis成功: PONG
	//Redis.Set(RedisCtx, "wxname", "路德维希", 0).Err()
	if err != nil {
		fmt.Println("Failed to set key:", err)
		return
	}
	return err
}
