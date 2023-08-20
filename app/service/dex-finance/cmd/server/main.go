package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	_ "go.uber.org/automaxprocs"
	"web3/app/service/dex-finance/internal/conf"
	"web3/app/service/dex-finance/internal/dao"
	"web3/app/service/dex-finance/internal/server"
	"web3/app/service/dex-finance/internal/service"
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
	grpcServer := server.NewGRPCServer(conf, svc, logger)
	httpServer := server.NewHTTPServer(conf, svc, logger)
	opts := app.GetAppOpts(app.APP_SERVICE_DEX_FINANCE)
	opts = append(opts,
		kratos.Server(
			grpcServer,
			httpServer,
		),
	)
	return kratos.New(opts...)
}

func main() {
	// default conf file path
	flagConf = "app/service/dex-finance/configs"
	flag.Parse()
	conf.Init(flagConf)
	bc := conf.Conf

	dao.New(bc)
	logger := app.GetAppLogger(app.APP_SERVICE_DEX_FINANCE)
	serverApp := newApp(bc, logger)
	if err := serverApp.Run(); err != nil {
		panic(err)
	}
}
