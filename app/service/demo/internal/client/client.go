package client

import (
	"github.com/go-kratos/kratos/v2/log"
	"web3/api/service/demo"
	dexFinance "web3/api/service/dex-finance"
	dexMember "web3/api/service/dex-member"
	"web3/app/service/demo/internal/conf"
	"web3/pkg/app"
)

var GlobalClient *Client

type Client struct {
	ClientAppServiceDemoUserGRPC       demo.UserClient
	ClientAppServiceDemoUserHTTP       demo.UserHTTPClient
	ClientAppServiceDexMemberUserGRPC  dexMember.UserClient
	ClientAppServiceDexMemberAuthGRPC  dexMember.AuthClient
	ClientAppServiceDexFinanceCoinGRPC dexFinance.CoinClient
}

func New(conf *conf.Bootstrap, logger log.Logger) *Client {
	c := &Client{}
	appServiceDemoHttpConn, err := app.GetHttpClientConn(app.APP_SERVICE_DEMO)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDemoUserHTTP = demo.NewUserHTTPClient(appServiceDemoHttpConn)

	appServiceDemoGrpcConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DEMO)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDemoUserGRPC = demo.NewUserClient(appServiceDemoGrpcConn)

	appServiceDexMemberGrpcConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DEX_MEMBER)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDexMemberUserGRPC = dexMember.NewUserClient(appServiceDexMemberGrpcConn)
	c.ClientAppServiceDexMemberAuthGRPC = dexMember.NewAuthClient(appServiceDexMemberGrpcConn)

	appServiceDexFinanceGrpcConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DEX_FINANCE)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDexFinanceCoinGRPC = dexFinance.NewCoinClient(appServiceDexFinanceGrpcConn)

	GlobalClient = c
	return c
}
