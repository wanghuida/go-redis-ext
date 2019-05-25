package model

import "time"

// UserStatus 用户状态的自定义常量
type UserStatus int

const (
	UserStatusValid   = 1   // UserStatusValid:   有效的
	UserStatusInvalid = 2   // UserStatusInvalid: 无效的
	UserStatusCheat   = 3   // UserStatusCheat:   作弊的
)

// UserInfo 用户其他信息
type UserInfo struct {
	Id       int64
	Nickname string
}

// User 用户基本信息
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
