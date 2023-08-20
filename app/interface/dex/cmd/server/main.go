package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	_ "go.uber.org/automaxprocs"
	"web3/app/interface/dex/internal/client"
	"web3/app/interface/dex/internal/conf"
	"web3/app/interface/dex/internal/server"
	"web3/app/interface/dex/internal/service"
	"web3/pkg/app"
)

var (
	// flagConf is the config flag.
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(conf *conf.Bootstrap, logger log.Logger) *kratos.App {
	svc := service.New(conf)
	httpServer := server.NewHTTPServer(conf, svc, logger)
	opts := app.GetAppOpts(app.APP_INTERFACE_DEX)
	opts = append(opts,
		kratos.Server(
			httpServer,
		),
	)
	return kratos.New(opts...)
}

func main() {
	// default conf file path
	flagConf = "app/interface/dex/configs"
	flag.Parse()
	conf.Init(flagConf)
	bc := conf.Conf
	client.New()
	logger := app.GetAppLogger(app.APP_INTERFACE_DEX)
	serverApp := newApp(bc, logger)
	if err := serverApp.Run(); err != nil {
		panic(err)
	}
}
