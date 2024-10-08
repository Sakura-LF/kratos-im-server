package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

// User 用户表
type User struct {
	gorm.Model
	Pwd      string `gorm:"size:64" json:"pwd"`
	Nickname string `gorm:"size:32" json:"nickname"`
	Abstract string `gorm:"size:128"  json:"abstract"`
	Avatar   string `gorm:"size:256"  json:"avatar"`
	IP       string `gorm:"size:32" json:"ip"`
	Addr     string `gorm:"size:64" json:"addr"`
	Role     int    `json:"role"` //角色1 管理员2 普通用户
	OpenID   string `gorm:"size:64" json:"token"`
}

// UserRepo  is a Greater repo.
type UserRepo interface {
	// Create 创建用户
	Create(context.Context, *User) (*User, error)

	// Set  设置Token
	Set(context.Context, string, string, time.Duration) *redis.StatusCmd
	// Get  获取Token
	Get(context.Context, string) *redis.StringCmd
}

// UserUsecase  is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) CreateUser() {
	fmt.Println("Test")
	return
}
