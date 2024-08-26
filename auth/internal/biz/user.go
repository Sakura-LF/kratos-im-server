package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	Role     int    `json:"role"` //角色1管理员2普通用户
}

// GreeterRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
}

// GreeterUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", "sakura")
	return uc.repo.Save(ctx, g)
}
