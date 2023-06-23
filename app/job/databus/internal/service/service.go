package service

import (
	"web3/app/job/databus/internal/conf"
)

type Service struct {
}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	return svc
}
