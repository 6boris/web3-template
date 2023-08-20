package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	_ "go.uber.org/automaxprocs"
	"os"
	"web3/app/interface/web3/internal/client"
	"web3/app/interface/web3/internal/conf"
	"web3/app/interface/web3/internal/server"
	"web3/app/interface/web3/internal/service"
	"web3/pkg/app"
)

var (
	// flagConf is the config flag.
	flagConf    string
	hostName, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(conf *conf.Bootstrap, logger log.Logger) *kratos.App {
	svc := service.New(conf)
	httpServer := server.NewHTTPServer(conf, svc, logger)
	opts := app.GetAppOpts(app.APP_INTERFACE_WEB3)
	opts = append(opts,
		kratos.Server(
			httpServer,
		),
	)
	return kratos.New(opts...)
}

func main() {
	// default conf file path
	flagConf = "app/interface/web3/configs"
	flag.Parse()
	conf.Init(flagConf)
	bc := conf.Conf
	logger := app.GetAppLogger(app.APP_INTERFACE_WEB3)
	client.New(conf.Conf, logger)
	serverApp := newApp(bc, logger)
	if err := serverApp.Run(); err != nil {
		panic(err)
	}
}
