package task

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"time"
)

type ServerOption func(o *Server)

type Server struct {
	CronJob *cron.Cron
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{}
	srv.CronJob = cron.New()
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	log.Context(ctx).Info("[Task] server starting")
	s.CronJob.Start()
	_ = s._startWatchTask(ctx)
	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	log.Context(ctx).Info("[Task] server stopping")
	s.CronJob.Stop()
	_ = s._stopWatchTask(ctx)
	time.Sleep(time.Second * 1)
	return nil
}
