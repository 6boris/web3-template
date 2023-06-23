package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"web3/api/service/demo"
	"web3/app/service/demo/internal/conf"
	"web3/pkg/otel"
)

var GlobalClient *Client

type Client struct {
	ClientAppServiceDemoUserGRPC demo.UserClient
	ClientAppServiceDemoUserHTTP demo.UserHTTPClient
}

func New(conf *conf.Bootstrap, logger log.Logger) *Client {
	clientConf := conf.GetClient()
	ctx := context.Background()
	c := &Client{}
	conn1, err := http.NewClient(context.Background(),
		http.WithEndpoint(clientConf.GetAppServiceDemoHttp()),
		http.WithMiddleware(
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

	c.ClientAppServiceDemoUserHTTP = demo.NewUserHTTPClient(conn1)

	conn2, err := grpc.DialInsecure(ctx,
		grpc.WithEndpoint(clientConf.GetAppServiceDemoGrpc()),
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
	c.ClientAppServiceDemoUserGRPC = demo.NewUserClient(conn2)
	GlobalClient = c
	return c
}
