package worker

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ServerOption func(o *Server)

type Server struct{}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{}
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	log.Context(ctx).Info("[Worker] server starting")
	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	log.Context(ctx).Info("[Worker] server stopping")
	return nil
}
