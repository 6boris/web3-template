package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"web3/app/service/dex-member/internal/conf"
	"web3/pkg/metrics"
)

var GlobalDao *Dao

// Dao ...
type Dao struct {
	MySQLClient *gorm.DB
	RedisClient *redis.Client
}

// New ...
func New(conf *conf.Bootstrap) *Dao {
	d := &Dao{}
	var err error
	d.MySQLClient, err = initMySQL(conf)
	if err != nil {
		panic(err)
	}
	d.RedisClient = initRedis(conf.GetApp(), conf.GetData().GetRedis())
	if err != nil {
		panic(err)
	}
	GlobalDao = d
	return d
}

// initRedis ...
func initRedis(appConf *conf.App, redisConf *conf.Redis) *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:         redisConf.GetAddr(),
		Password:     redisConf.GetPassword(),
		DB:           int(redisConf.GetDb()),
		ReadTimeout:  redisConf.GetReadTimeout().AsDuration(),
		WriteTimeout: redisConf.GetWriteTimeout().AsDuration(),
	})
	RedisClient.AddHook(metrics.RedisMetricsHook{
		Name:            redisConf.GetName(),
		ServiceID:       appConf.GetId(),
		ServiceInstance: appConf.GetInstance(),
	})
	return RedisClient
}

func initMySQL(conf *conf.Bootstrap) (*gorm.DB, error) {
	var err error
	client, err := gorm.Open(mysql.New(mysql.Config{
		DSN: conf.GetData().GetMysql().GetDsn(),
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})
	if err != nil {
		return nil, err
	}
	err = client.Use(&metrics.GormMetricsPlugin{
		InstName:        conf.GetData().GetMysql().GetName(),
		ServiceID:       conf.GetApp().GetId(),
		ServiceInstance: conf.GetApp().GetInstance(),
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (d *Dao) Ping(ctx context.Context) error {
	var err error
	err = d.RedisClient.Ping(ctx).Err()
	if err != nil {
		return err
	}
	return err
}
