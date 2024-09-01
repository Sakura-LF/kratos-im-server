package biz

import (
	"context"
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
}

// UserRepo  is a Greater repo.
type UserRepo interface {
	// FindByUserName 1df1
	FindByUserName(context.Context, string) (*User, error)

	// Set  设置Token
	Set(context.Context, string, string, time.Duration) *redis.StatusCmd
	// Get
	Get(context.Context, string) *redis.StringCmd
}

// UserUsecase  is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// GetUser CreateUser CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) GetUser(ctx context.Context, userName string) (*User, error) {
	return uc.repo.FindByUserName(ctx, userName)
}

func (uc *UserUsecase) SetToken(ctx context.Context, key, token string, expiration time.Duration) error {
	return uc.repo.Set(ctx, key, token, expiration).Err()
}
func (uc *UserUsecase) GetToken(ctx context.Context, key string) (string, error) {
	return uc.repo.Get(ctx, key).Result()
}
