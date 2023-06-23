package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"web3/app/job/datawatch/internal/conf"
	"web3/app/job/datawatch/internal/service"
	"web3/app/job/datawatch/internal/task"
)

func NewTASKServer(c *conf.Bootstrap, s *service.Service, logger log.Logger) *task.Server {
	var opts []task.ServerOption
	srv := task.NewServer(opts...)
	return srv
}
