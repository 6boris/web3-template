package worker

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type ServerOption func(o *Server)

type Server struct {
	CronJob *cron.Cron
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{}
	srv.CronJob = cron.New()
	_, _ = srv.CronJob.AddFunc("@every 2s", srv._demoClient)
	_, _ = srv.CronJob.AddFunc("@every 10s", srv._demoDao)
	_, _ = srv.CronJob.AddFunc("@every 2s", srv._demoDefiPrice)

	return srv
}

func (s *Server) Start(ctx context.Context) error {
	log.Context(ctx).Info("[Worker] server starting")
	s.CronJob.Start()
	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	log.Context(ctx).Info("[Worker] server stopping")
	s.CronJob.Stop()
	return nil
}
