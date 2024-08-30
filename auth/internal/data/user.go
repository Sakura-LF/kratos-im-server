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

func (u *userRepo) FindByUserName(ctx context.Context, userName string) (*biz.User, error) {
	var user *biz.User
	if err := u.data.db.Take(&user, "id = ?", userName).Error; err != nil {
		u.log.Error(err)
		return nil, err
	}
	return user, nil
}

// NewUserRepo  .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
