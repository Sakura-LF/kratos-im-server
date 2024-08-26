package data

import (
	"auth/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) FindByID(ctx context.Context, i int64) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) ListByHello(ctx context.Context, s string) ([]*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

// NewUserRepo  .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
