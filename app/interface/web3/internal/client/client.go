package client

import (
	"github.com/go-kratos/kratos/v2/log"
	"web3/api/service/data"
	"web3/api/service/demo"
	"web3/app/interface/web3/internal/conf"
	"web3/pkg/app"
)

var GlobalClient *Client

type Client struct {
	ClientAppServiceDemoUserGRPC demo.UserClient
	ClientAppServiceDataDefiGRPC data.DefiClient
}

func New(conf *conf.Bootstrap, logger log.Logger) *Client {
	c := &Client{}
	appServiceDemoConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DEMO)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDemoUserGRPC = demo.NewUserClient(appServiceDemoConn)

	appServiceDataConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DATA)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDataDefiGRPC = data.NewDefiClient(appServiceDataConn)

	GlobalClient = c
	return c
}
