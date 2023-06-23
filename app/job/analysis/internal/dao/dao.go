package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"web3/app/job/analysis/internal/conf"
	"web3/pkg/metrics"
)

var GlobalDao *Dao

type Dao struct {
	Web3MySQLClient *gorm.DB
}

func New(conf *conf.Bootstrap) *Dao {
	d := &Dao{}

	client, err := gorm.Open(mysql.New(mysql.Config{
		DSN: conf.GetData().GetMysql().GetDsn(),
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})
	if err != nil {
		panic(err)
	}
	err = client.Use(&metrics.GormMetricsPlugin{
		InstName:        conf.GetData().GetMysql().GetName(),
		ServiceID:       conf.GetApp().GetId(),
		ServiceInstance: conf.GetApp().GetInstance(),
	})
	if err != nil {
		panic(err)
	}
	d.Web3MySQLClient = client
	GlobalDao = d
	return d
}
func (d *Dao) Close() error {
	return nil
}
