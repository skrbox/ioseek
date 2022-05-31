package model

import (
	"time"
)

// 友链
type Link struct {
	URL          string `gorm:"unique"`
	Name         string `gorm:"unique"`
	IsVerified   bool
	IsSubscribed bool // 是否订阅
	meta
}

// 账户
type Profile struct {
	Nickname  string    `gorm:"unique"` // 用户昵称
	Email     string    `gorm:"unique"`
	IsOwner   bool      // 是否超管
	IsActive  bool      // 是否有效
	Password  string    // 绑定口令
	ExpiredAt time.Time // 口令有效期，到期未绑定则作废
	meta
}

// 告警通知对象
const (
	owner  = "owner"
	admin  = "admin"
	author = "author"
)

// 告警通知
type Event struct {
	Event  string `json:"event"`
	Detail string // 事件详情
	meta
}
