package client

import (
	"github.com/go-kratos/kratos/v2/log"
	"web3/app/service/data/internal/conf"
)

var GlobalClient *Client

type Client struct {
}

func New(conf *conf.Bootstrap, logger log.Logger) *Client {
	c := &Client{}
	GlobalClient = c
	return c
}
