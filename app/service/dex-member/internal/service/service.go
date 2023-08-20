package service

import (
	pb "web3/api/service/dex-member"
	"web3/app/service/dex-member/internal/conf"
	"web3/app/service/dex-member/internal/dao"
)

type Service struct {
	pb.UnimplementedUserServer
	pb.UnimplementedAuthServer

	dao *dao.Dao
}

func New(conf *conf.Bootstrap) *Service {
	svc := &Service{}
	svc.dao = dao.New(conf)
	return svc
}
