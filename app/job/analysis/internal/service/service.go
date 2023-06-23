package service

import (
	"web3/app/job/analysis/internal/conf"
)

type Service struct {
}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	return svc
}
