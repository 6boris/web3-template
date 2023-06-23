package service

import (
	pb "web3/api/service/demo"
	"web3/app/service/demo/internal/conf"
	"web3/app/service/demo/internal/dao"
)

type Service struct {
	pb.UnimplementedUserServer

	dao *dao.Dao
}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	svc.dao = dao.New(conf)
	return svc
}
