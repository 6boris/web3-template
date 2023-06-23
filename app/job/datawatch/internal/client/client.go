package client

import (
	web3Client "github.com/6boris/web3-go/client"
	"github.com/go-kratos/kratos/v2/log"
	"web3/app/job/datawatch/internal/conf"
)

var GlobalClient *Client

type Client struct {
	Web3ClientPool *web3Client.Pool
}

func New(conf *conf.Bootstrap, logger log.Logger) *Client {
	c := &Client{}
	clientPoolConf := web3Client.GetDefaultConfPool()
	clientPoolConf.AppID = "app.job.databus"
	c.Web3ClientPool = web3Client.NewPool(clientPoolConf)
	GlobalClient = c
	return c
}
