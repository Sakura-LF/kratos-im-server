package data

import (
	"auth/internal/biz"
	"auth/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	log2 "log"
	"os"
	"time"
)

// DataProviderSet  is data providers.
var DataProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewDB .
func NewDB(c *conf.Data, zlog log.Logger) *gorm.DB {

	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             300 * time.Millisecond, // 慢查询 SQL 阈值
			Colorful:                  true,                   // 是否启动彩色打印
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Error, // Log lever
		},
	)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	zerolog := log.NewHelper(zlog)
	if err != nil {
		zerolog.Error(err)
		panic("failed to connect database")
	} else {
		zerolog.Info("mysql connect success")
	}
	return db
}

// NewRedis 创建 Redis 连接
func NewRedis(c *conf.Data, zlog log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Username:     c.Redis.Username, // default user
		Password:     c.Redis.Password,
		WriteTimeout: time.Duration(c.Redis.WriteTimeout.Seconds),
		ReadTimeout:  time.Duration(c.Redis.ReadTimeout.Seconds),
		DB:           0,               // use default DB
		DialTimeout:  1 * time.Second, // 1 second
	})

	zerolog := log.NewHelper(zlog)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		zerolog.Error(err)
		panic("Redis 连接失败")
	} else {
		zerolog.Info("Redis connect success")
	}
	return rdb
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	if err := db.AutoMigrate(&biz.User{}); err != nil {
		log.NewHelper(logger).Info("user table migrate fail")
	} else {
		log.NewHelper(logger).Info("user table migrate success")
	}
	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}
