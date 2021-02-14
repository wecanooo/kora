package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/wecanooo/kora/core"
)

// SetupRedis creates and returns a redis client instance
func SetupRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: 	core.GetConfig().DefaultString("REDIS.ADDR", "localhost"),
		Password: core.GetConfig().DefaultString("REDIS.PASSWORD", ""),
		DB: core.GetConfig().DefaultInt("REDIS.DATABASE", 0),
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic("Redis 연결설정을 하는 중 오류가 발생되었습니다, 오류: " + err.Error())
	} else {
		fmt.Printf("redis connect ping responses: %s", pong)
	}

	return client
}