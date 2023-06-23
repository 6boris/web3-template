package service

import (
	"web3/app/interface/web3/internal/conf"
	"web3/app/interface/web3/internal/service/defi"
	"web3/app/interface/web3/internal/service/demo"
)

type Service struct {
	demo.DemoService
	defi.DefiService
	//pb.UnsafeDemoServer
}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	return svc
}
