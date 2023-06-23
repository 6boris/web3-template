package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"web3/app/service/demo/internal/conf"
	"web3/app/service/demo/internal/service"
	"web3/app/service/demo/internal/worker"
)

func NewWORKERServer(c *conf.Bootstrap, s *service.Service, logger log.Logger) *worker.Server {
	var opts []worker.ServerOption
	srv := worker.NewServer(opts...)
	return srv
}
