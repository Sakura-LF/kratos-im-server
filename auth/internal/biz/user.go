package biz

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Pwd      string `gorm:"size:64" json:"pwd"`
	Nickname string `gorm:"size:32" json:"nickname"`
	Abstract string `gorm:"size:128"  json:"abstract"`
	Avatar   string `gorm:"size:256"  json:"avatar"`
	IP       string `gorm:"size:32" json:"ip"`
	Addr     string `gorm:"size:64" json:"addr"`
	Role     int    `json:"role"` //角色1管理员2普通用户
}
