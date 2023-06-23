package consumer

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ServerOption func(o *Server)

type Server struct {
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{}
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	log.Context(ctx).Info("[Consumer] server starting")
	_ = s._startConsumers(ctx)
	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	log.Context(ctx).Info("[Consumer] server stopping")
	_ = s._stopConsumers(ctx)
	return nil
}

var InstList = []Consumer{
	NewBlockEth(),
	NewBlockAptos(),
}

func (s *Server) _startConsumers(ctx context.Context) error {
	for _, v := range InstList {
		loopErr := v.Start(ctx)
		if loopErr != nil {
			log.Context(ctx).Errorf("[Consumer] Start Name:%+v err:%+v", v.Name(), loopErr)
		}
	}
	return nil
}

func (s *Server) _stopConsumers(ctx context.Context) error {
	for _, v := range InstList {
		loopErr := v.Stop(ctx)
		if loopErr != nil {
			log.Context(ctx).Errorf("[Consumer] Stop Name:%+v err:%+v", v.Name(), loopErr)
		}
	}
	return nil
}

type Consumer interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Reload(ctx context.Context) error
}
