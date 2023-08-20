package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"web3/api/service/data"
	"web3/api/service/demo"
	"web3/app/interface/web3/internal/conf"
	"web3/pkg/app"
	"web3/pkg/otel"
)

var GlobalClient *Client

type Client struct {
	ClientAppServiceDemoUserGRPC demo.UserClient
	ClientAppServiceDataDefiGRPC data.DefiClient
}

func New(conf *conf.Bootstrap, logger log.Logger) *Client {
	clientConf := conf.GetClient()
	ctx := context.Background()
	c := &Client{}

	//cli, err := clientv3.New(clientv3.Config{
	//	Endpoints: []string{"127.0.0.1:2379"},
	//})
	//if err != nil {
	//	panic(err)
	//}

	//conn1, err := http.NewClient(
	//	context.Background(),
	//	http.WithEndpoint("discovery:///app.service.demo"),
	//	http.WithDiscovery(etcd.New(cli)),
	//	http.WithBlock(),
	//	http.WithMiddleware(
	//		recovery.Recovery(),
	//		tracing.Client(),
	//		logging.Client(logger),
	//		otel.MetricsClient(&otel.Conf{
	//			AppID:    conf.GetApp().GetId(),
	//			Env:      conf.GetApp().GetEnv(),
	//			Instance: conf.GetApp().GetInstance(),
	//			Cluster:  conf.GetApp().GetCluster(),
	//			Zone:     conf.GetApp().GetZone(),
	//			Version:  conf.GetApp().GetVersion(),
	//		}),
	//	),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//conn1, err := http.NewClient(context.Background(),
	//	http.WithEndpoint(clientConf.GetAppServiceDemoHttp()),
	//	http.WithMiddleware(
	//		recovery.Recovery(),
	//		tracing.Client(),
	//		logging.Client(logger),
	//		otel.MetricsClient(&otel.Conf{
	//			AppID:    conf.GetApp().GetId(),
	//			Env:      conf.GetApp().GetEnv(),
	//			Instance: conf.GetApp().GetInstance(),
	//			Cluster:  conf.GetApp().GetCluster(),
	//			Zone:     conf.GetApp().GetZone(),
	//			Version:  conf.GetApp().GetVersion(),
	//		}),
	//	),
	//)
	//if err != nil {
	//	panic(err)
	//}

	appServiceDemoConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DEMO)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDemoUserGRPC = demo.NewUserClient(appServiceDemoConn)
	conn3, err := grpc.DialInsecure(ctx,
		grpc.WithEndpoint(clientConf.GetAppServiceDataGrpc()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			logging.Client(logger),
			otel.MetricsClient(&otel.Conf{
				AppID:    conf.GetApp().GetId(),
				Env:      conf.GetApp().GetEnv(),
				Instance: conf.GetApp().GetInstance(),
				Cluster:  conf.GetApp().GetCluster(),
				Zone:     conf.GetApp().GetZone(),
				Version:  conf.GetApp().GetVersion(),
			}),
		),
	)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDataDefiGRPC = data.NewDefiClient(conn3)
	GlobalClient = c
	return c
}
