package service

import (
	"web3/app/interface/dex/internal/conf"
)

type Service struct{}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	return svc
}
