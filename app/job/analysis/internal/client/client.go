package client

import (
	"context"
	web3Client "github.com/6boris/web3-go/client"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"web3/api/service/data"
	"web3/app/job/analysis/internal/conf"
	"web3/pkg/otel"
)

var GlobalClient *Client

type Client struct {
	Web3ClientPool               *web3Client.Pool
	ClientAppServiceDataChainRPC data.ChainClient
	ClientAppServiceDataDefiGRPC data.DefiClient
}

func New(conf *conf.Bootstrap, logger log.Logger) *Client {
	c := &Client{}
	clientPoolConf := web3Client.GetDefaultConfPool()
	clientPoolConf.AppID = "app.job.analysis"
	c.Web3ClientPool = web3Client.NewPool(clientPoolConf)
	ctx := context.Background()
	clientConf := conf.GetClient()
	conn1, err := grpc.DialInsecure(ctx,
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
	c.ClientAppServiceDataChainRPC = data.NewChainClient(conn1)
	c.ClientAppServiceDataDefiGRPC = data.NewDefiClient(conn1)

	GlobalClient = c
	return c
}
