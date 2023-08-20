package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	_ "go.uber.org/automaxprocs"
	"web3/app/service/demo/internal/client"
	"web3/app/service/demo/internal/conf"
	"web3/app/service/demo/internal/dao"
	"web3/app/service/demo/internal/server"
	"web3/app/service/demo/internal/service"
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
	workerServer := server.NewWORKERServer(conf, svc, logger)
	opts := app.GetAppOpts(app.APP_SERVICE_DEMO)
	opts = append(opts,
		kratos.Server(
			grpcServer,
			httpServer,
			workerServer,
		))
	return kratos.New(opts...)
}

func main() {
	// default conf file path
	flagConf = "app/service/demo/configs"
	flag.Parse()
	conf.Init(flagConf)
	bc := conf.Conf
	dao.New(bc)
	logger := app.GetAppLogger(app.APP_SERVICE_DEMO)
	client.New(bc, logger)
	serverApp := newApp(bc, logger)
	if err := serverApp.Run(); err != nil {
		panic(err)
	}
}
