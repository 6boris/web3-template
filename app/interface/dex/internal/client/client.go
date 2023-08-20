package client

import (
	dexFinance "web3/api/service/dex-finance"
	dexMember "web3/api/service/dex-member"
	"web3/pkg/app"
)

var GlobalClient *Client

type Client struct {
	ClientAppServiceDexMemberUserGRPC     dexMember.UserClient
	ClientAppServiceDexMemberAuthGRPC     dexMember.AuthClient
	ClientAppServiceDexFinanceAccountGRPC dexFinance.AccountClient
}

func New() *Client {
	c := &Client{}
	appServiceDexMemberGrpcConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DEX_MEMBER)
	if err != nil {
		panic(err)
	}
	appServiceDexFinanceGrpcConn, err := app.GetGrpcClientConn(app.APP_SERVICE_DEX_FINANCE)
	if err != nil {
		panic(err)
	}
	c.ClientAppServiceDexMemberUserGRPC = dexMember.NewUserClient(appServiceDexMemberGrpcConn)
	c.ClientAppServiceDexMemberAuthGRPC = dexMember.NewAuthClient(appServiceDexMemberGrpcConn)
	c.ClientAppServiceDexFinanceAccountGRPC = dexFinance.NewAccountClient(appServiceDexFinanceGrpcConn)
	GlobalClient = c
	return c
}
