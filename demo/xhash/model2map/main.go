package main

import (
	"github.com/go-redis/redis"
	"github.com/k0kubun/pp"
	"github.com/wanghuida/go-redis-ext/demo/model"
	"github.com/wanghuida/go-redis-ext/xredis/xhash"
	"time"
)

func main() {

	// 数据的准备
	redPacketId := new(int64)
	*redPacketId = 10001

	bob := model.UserInfo{Id: 2, Nickname: "Bob"}
	wade := model.UserInfo{Id: 3, Nickname: "Wade"}

	now := time.Now().Local()

	user := &model.User{
		Id:          1,
		RedPacketId: redPacketId,
		Name:        "william",
		Tags:        []string{"man", "pupil"},
		Status:      model.UserStatusValid,
		IsNew:       true,
		Score:       3.1415,
		Friends:     map[int64]model.UserInfo{bob.Id: bob, wade.Id: wade},
		Info:        &bob,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}

	// struct 转 hash 可以使用的 map
	result, err := xhash.Model2map(user)
	if err != nil {
		panic(err)
	}
	pp.Println(result)

	// 存储数据
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	err = redisClient.HMSet("user1", result).Err()
	if err != nil {
		panic(err)
	}
}
