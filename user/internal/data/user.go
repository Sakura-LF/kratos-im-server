package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"time"
	"user/internal/biz"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func (u *UserRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Set(ctx context.Context, s string, s2 string, duration time.Duration) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Get(ctx context.Context, s string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

// NewUserRepo  .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
