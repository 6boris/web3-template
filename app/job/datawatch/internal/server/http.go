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
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/propagation"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	"web3/app/job/datawatch/internal/conf"
	"web3/app/job/datawatch/internal/service"
	"web3/pkg/otel"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, s *service.Service, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		)),
		http.Middleware(
			recovery.Recovery(),
			metadata.Server(metadata.WithPropagatedPrefix("")),
			logging.Server(logger),
			tracing.Server(
				tracing.WithTracerProvider(traceSdk.NewTracerProvider()),
				tracing.WithPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})),
			),
			otel.MetricsServer(&otel.Conf{
				AppID:    c.GetApp().GetId(),
				Env:      c.GetApp().GetEnv(),
				Instance: c.GetApp().GetInstance(),
				Cluster:  c.GetApp().GetCluster(),
				Zone:     c.GetApp().GetZone(),
				Version:  c.GetApp().GetVersion(),
			}),
			ratelimit.Server(ratelimit.WithLimiter(bbr.NewLimiter())),
			validate.Validator(),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	return srv
}
