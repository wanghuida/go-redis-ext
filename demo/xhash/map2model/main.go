package main

import (
	"github.com/go-redis/redis"
	"github.com/k0kubun/pp"
	"github.com/wanghuida/go-redis-ext/demo/model"
	"github.com/wanghuida/go-redis-ext/xredis/xhash"
)

func main() {

	// 读取数据
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	result := redisClient.HGetAll("user1").Val()

	user := new(model.User)
	err := xhash.Map2model(result, user)
	if err != nil {
		panic(err)
	}

	pp.Println(user)

}
