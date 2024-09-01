package data

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func (u *userRepo) Set(ctx context.Context, key string, value string, duration time.Duration) *redis.StatusCmd {
	return u.data.rdb.Set(ctx, key, value, duration)
}

func (u *userRepo) Get(ctx context.Context, s string) *redis.StringCmd {
	return u.data.rdb.Get(ctx, s)
}
