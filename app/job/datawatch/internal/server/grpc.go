package server

import (
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.opentelemetry.io/otel/propagation"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	"web3/app/job/datawatch/internal/conf"
	"web3/app/job/datawatch/internal/service"
	"web3/pkg/otel"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, s *service.Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(metadata.WithPropagatedPrefix("")),
			logging.Server(logger),
			tracing.Server(
				tracing.WithTracerProvider(traceSdk.NewTracerProvider()),
				tracing.WithPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})),
			),
			ratelimit.Server(ratelimit.WithLimiter(bbr.NewLimiter())),
			validate.Validator(),
			otel.MetricsServer(&otel.Conf{
				AppID:    c.GetApp().GetId(),
				Env:      c.GetApp().GetEnv(),
				Instance: c.GetApp().GetInstance(),
				Cluster:  c.GetApp().GetCluster(),
				Zone:     c.GetApp().GetZone(),
				Version:  c.GetApp().GetVersion(),
			}),
		),
	}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	return srv
}
