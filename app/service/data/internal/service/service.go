package service

import (
	"web3/app/service/data/internal/conf"
	"web3/app/service/data/internal/service/chain"
	"web3/app/service/data/internal/service/defi"
)

type Service struct {
	chain.ChainService
	defi.DefiService
}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	return svc
}
