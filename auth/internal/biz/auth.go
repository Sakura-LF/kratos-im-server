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
	OpenID   string `gorm:"size:64" json:"token"`
}

// AuthRepo  is a Greater repo.
type AuthRepo interface {
	// FindByUserName 1df1
	FindByUserName(context.Context, string) (*User, error)

	// Set  设置Token
	Set(context.Context, string, string, time.Duration) *redis.StatusCmd
	// Get
	Get(context.Context, string) *redis.StringCmd
}

// AuthUsecase  is a User usecase.
type AuthUsecase struct {
	repo AuthRepo
	log  *log.Helper
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// GetUser CreateUser CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *AuthUsecase) GetUser(ctx context.Context, userName string) (*User, error) {
	return uc.repo.FindByUserName(ctx, userName)
}

func (uc *AuthUsecase) SetToken(ctx context.Context, key, token string, expiration time.Duration) error {
	return uc.repo.Set(ctx, key, token, expiration).Err()
}
func (uc *AuthUsecase) GetToken(ctx context.Context, key string) (string, error) {
	return uc.repo.Get(ctx, key).Result()
}
