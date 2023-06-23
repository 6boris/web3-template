package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	_ "go.uber.org/automaxprocs"
	"os"
	"web3/app/service/demo/internal/client"
	"web3/app/service/demo/internal/conf"
	"web3/app/service/demo/internal/dao"
	"web3/app/service/demo/internal/server"
	"web3/app/service/demo/internal/service"
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
	grpcServer := server.NewGRPCServer(conf, svc, logger)
	httpServer := server.NewHTTPServer(conf, svc, logger)
	workerServer := server.NewWORKERServer(conf, svc, logger)

	return kratos.New(
		kratos.ID(conf.GetApp().GetId()),
		kratos.Name(conf.GetApp().GetId()),
		kratos.Version(conf.GetApp().GetVersion()),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			grpcServer,
			httpServer,
			workerServer,
		),
	)
}

func main() {
	// default conf file path
	flagConf = "app/service/demo/configs"
	flag.Parse()
	conf.Init(flagConf)
	bc := conf.Conf

	dao.New(bc)

	logger := log.With(
		log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"source", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
		"service.id", bc.GetApp().GetId(),
		"service.env", bc.GetApp().GetEnv(),
		"service.instance", hostName,
		"service.cluster", bc.GetApp().GetCluster(),
		"service.zone", bc.GetApp().GetZone(),
		"service.version", bc.GetApp().GetVersion(),
	)
	client.New(bc, logger)
	app := newApp(bc, logger)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
