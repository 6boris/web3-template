package service

import (
	pb "web3/api/service/dex-finance"
	"web3/app/service/dex-finance/internal/conf"
)

type Service struct {
	pb.UnimplementedCoinServer
	pb.UnimplementedAccountServer
}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	return svc
}
