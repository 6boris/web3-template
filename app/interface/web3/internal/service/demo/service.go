package demo

import (
	pb "web3/api/interface/web3"
	"web3/app/interface/web3/internal/conf"
)

type DemoService struct {
	pb.UnimplementedDemoServer
}

func New(conf *conf.Bootstrap) *DemoService {
	svc := &DemoService{}
	return svc
}
