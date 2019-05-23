# 提高 go-redis 的 hash 易用性，提供 map 与 model 互相转换

## 安装

```shell
go get -u github.com/wanghuida/go-redis-ext
```

## 功能

- 读取 hash 类型数据可以方便的将 map 转成结构体
- 写入 hash 类型数据可以将结构体转成 map
- 结构体可以被其他 sdk 复用，例如 xorm, gorm 等

## 类型支持

- `Bool` `*Bool`
- `Int` `Int8` `Int16` `Int32` `Int64`
- `*Int` `*Int8` `*Int16` `*Int32` `*Int64`
- `Uint` `Uint8` `Uint16` `Uint32` `Uint64`
- `*Uint` `*Uint8` `*Uint16` `*Uint32` `*Uint64`
- `Float32` `*Float32` `Float64` `*Float64`
- `Interface`
- `Map`
- `Slice`
- `String` `*String`
- `Struct` `*Struct`
- `time.Time` `*time.Time`
- `alias` 例: `type UserStatus int`

## 快速开始

```go
// model 转 map
result, err := xhash.Model2map(yourModel)

// map 转 model
yourModel := new(Model)
err := xhash.Map2model(mapVar, yourModel)
```


## 案例

### 定义模型，以用户信息为例

```go
package model

import "time"

// 用户状态的自定义常量
type UserStatus int
const (
	UserStatusValid   = 1
	UserStatusInvalid = 2
	UserStatusCheat   = 3
)

// 用户其他信息
type UserInfo struct {
	Id       int64
	Nickname string
}

// 用户基本信息
type User struct {
	Id          int64
	RedPacketId *int64
	Name        string
	Tags        []string
	Status      UserStatus
	IsNew       bool
	Score       float64
	Friends     map[int64]UserInfo
	Info        *UserInfo `redis:"user_info"` // 自定义存储名称
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Ignore      string `redis:"-"` // 忽略该字段，redis 不存储
}

```

### 写入数据

```go
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

	bob := model.UserInfo{Id: 2, Nickname:"Bob"}
	wade := model.UserInfo{Id: 3, Nickname:"Wade"}

	now := time.Now().Local()

	user := &model.User{
		Id: 1,
		RedPacketId: redPacketId,
		Name: "william",
		Tags: []string{"man", "pupil"},
		Status: model.UserStatusValid,
		IsNew: true,
		Score: 3.1415,
		Friends: map[int64]model.UserInfo{bob.Id: bob, wade.Id: wade},
		Info: &bob,
		CreatedAt: now,
		UpdatedAt: &now,
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

```

```shell
# 输出
127.0.0.1:6379> HGETALL user1
 1) "red_packet_id"
 2) "10001"
 3) "score"
 4) "3.1415"
 5) "user_info"
 6) "{\"Id\":2,\"Nickname\":\"Bob\"}"
 7) "updated_at"
 8) "2019-05-22 14:25:41"
 9) "created_at"
10) "2019-05-22 14:25:41"
11) "id"
12) "1"
13) "name"
14) "william"
15) "tags"
16) "[\"man\",\"pupil\"]"
17) "status"
18) "2"
19) "is_new"
20) "1"
21) "friends"
22) "{\"2\":{\"Id\":2,\"Nickname\":\"Bob\"},\"3\":{\"Id\":3,\"Nickname\":\"Wade\"}}"
```

### 读取数据

```go
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
```


```shell
# 输出
&model.User{
  Id:          1,
  RedPacketId: &10001,
  Name:        "william",
  Tags:        []string{
    "man",
    "pupil",
  },
  Status:  2,
  IsNew:   true,
  Score:   3.141500,
  Friends: map[int64]model.UserInfo{
    2: model.UserInfo{
      Id:       2,
      Nickname: "Bob",
    },
    3: model.UserInfo{
      Id:       3,
      Nickname: "Wade",
    },
  },
  Info: &model.UserInfo{
    Id:       2,
    Nickname: "Bob",
  },
  CreatedAt: 2019-05-22 14:25:41 Local,
  UpdatedAt: &2019-05-22 14:25:41 Local,
  Ignore:    "",
}
```

