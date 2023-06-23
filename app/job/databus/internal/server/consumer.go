package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"web3/app/job/databus/internal/conf"
	"web3/app/job/databus/internal/consumer"
	"web3/app/job/databus/internal/service"
)

func NewConsumerServer(c *conf.Bootstrap, s *service.Service, logger log.Logger) *consumer.Server {
	var opts []consumer.ServerOption
	srv := consumer.NewServer(opts...)
	return srv
}
