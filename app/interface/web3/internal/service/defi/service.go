package defi

import (
	pb "web3/api/interface/web3"
	"web3/app/interface/web3/internal/conf"
)

type DefiService struct {
	pb.UnimplementedDefiServer
}

func New(conf *conf.Bootstrap) *DefiService {
	svc := &DefiService{}
	return svc
}
