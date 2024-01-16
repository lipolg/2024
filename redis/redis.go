package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.136.56.76:6380",
		Password: "321084Li", // 没有密码，默认值
		DB:       0,          // 默认DB 0
	})
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Println(err)
	}
	rdb.Set(ctx, "id", 1, time.Hour)
	get := rdb.Get(ctx, "id")
	fmt.Println(get)
}
