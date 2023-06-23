package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

var Conf *Bootstrap

func Init(flagConf string) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	Conf = &bc
}
