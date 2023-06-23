package otel

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"gorm.io/gorm"
	"time"
)

var (
	ServerRequestCounter       api.Int64Counter
	ServerRequestHistogram     api.Int64Histogram
	ClientRequestCounter       api.Int64Counter
	ClientRequestHistogram     api.Int64Histogram
	MetricsMySQLRequestCounter api.Int64Counter
)

func init() {
	exporter, err := prometheus.New(prometheus.WithoutTargetInfo())
	if err != nil {
		log.Fatal(err)
	}
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter("metrics", api.WithInstrumentationVersion(runtime.Version()))
	counter, err := meter.Int64Counter("server_request", api.WithDescription("MetricsServer receive request counter"))
	if err != nil {
		log.Fatal(err)
	}
	his, err := meter.Int64Histogram("server_request", api.WithDescription("MetricsServer receive request histogram"))
	ServerRequestCounter = counter
	ServerRequestHistogram = his

	cCounter, err := meter.Int64Counter("client_request", api.WithDescription("MetricsClient receive request counter"))
	if err != nil {
		log.Fatal(err)
	}
	cHis, err := meter.Int64Histogram("client_request", api.WithDescription("MetricsClient receive request histogram"))
	ClientRequestCounter = cCounter
	ClientRequestHistogram = cHis

	c2, err := meter.Int64Counter("mysql_request", api.WithDescription("MySQL receive request counter"))
	ServerRequestCounter = counter
	ServerRequestHistogram = his
	MetricsMySQLRequestCounter = c2

}

func MetricsServer(conf *Conf) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			var (
				code      int
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}

			reply, err := handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}
			opt := api.WithAttributes(
				attribute.Key("service_id").String(conf.AppID),
				attribute.Key("service_version").String(conf.Version),
				attribute.Key("env").String(conf.Env),
				attribute.Key("instance").String(conf.Instance),
				attribute.Key("kind").String(kind),
				attribute.Key("operation").String(operation),
				attribute.Key("status_code").Int(code),
				attribute.Key("reason").String(reason),
			)
			ServerRequestCounter.Add(ctx, 1, opt)
			ServerRequestHistogram.Record(ctx, time.Since(startTime).Milliseconds(), opt)
			return reply, err
		}
	}
}
func MetricsClient(conf *Conf) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				code      int
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err := handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}

			opt := api.WithAttributes(
				attribute.Key("service_id").String(conf.AppID),
				attribute.Key("service_version").String(conf.Version),
				attribute.Key("env").String(conf.Env),
				attribute.Key("instance").String(conf.Instance),
				attribute.Key("kind").String(kind),
				attribute.Key("operation").String(operation),
				attribute.Key("status_code").Int(code),
				attribute.Key("reason").String(reason),
			)
			ClientRequestCounter.Add(ctx, 1, opt)
			ClientRequestHistogram.Record(ctx, time.Since(startTime).Milliseconds(), opt)
			return reply, err
		}
	}
}

type MetricsGormPlugin struct {
	DatabaseName string
	Conf         Conf
}

func (p *MetricsGormPlugin) Name() string {
	return ""
}
func (p *MetricsGormPlugin) Initialize(db *gorm.DB) error {
	_ = db.Callback().Query().After("gorm:query").Register("after:query", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		opt := api.WithAttributes(
			attribute.Key("app_id").String(p.Conf.AppID),
			attribute.Key("app_version").String(p.Conf.Version),
			attribute.Key("env").String(p.Conf.Env),
			attribute.Key("instance").String(p.Conf.Instance),
			attribute.Key("database_name").String(p.DatabaseName),
			attribute.Key("table").String(db.Statement.Table),
			attribute.Key("cmd").String(cmd),
		)
		MetricsMySQLRequestCounter.Add(context.TODO(), 1, opt)
	})
	_ = db.Callback().Create().After("gorm:create").Register("after:create", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		opt := api.WithAttributes(
			attribute.Key("app_id").String(p.Conf.AppID),
			attribute.Key("app_version").String(p.Conf.Version),
			attribute.Key("env").String(p.Conf.Env),
			attribute.Key("instance").String(p.Conf.Instance),
			attribute.Key("database_name").String(p.DatabaseName),
			attribute.Key("table").String(db.Statement.Table),
			attribute.Key("cmd").String(cmd),
		)
		MetricsMySQLRequestCounter.Add(context.TODO(), 1, opt)
	})
	_ = db.Callback().Update().After("gorm:update").Register("after:update", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		opt := api.WithAttributes(
			attribute.Key("app_id").String(p.Conf.AppID),
			attribute.Key("app_version").String(p.Conf.Version),
			attribute.Key("env").String(p.Conf.Env),
			attribute.Key("instance").String(p.Conf.Instance),
			attribute.Key("database_name").String(p.DatabaseName),
			attribute.Key("table").String(db.Statement.Table),
			attribute.Key("cmd").String(cmd),
		)
		MetricsMySQLRequestCounter.Add(context.TODO(), 1, opt)
	})
	_ = db.Callback().Delete().After("gorm:delete").Register("after:delete", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		opt := api.WithAttributes(
			attribute.Key("app_id").String(p.Conf.AppID),
			attribute.Key("app_version").String(p.Conf.Version),
			attribute.Key("env").String(p.Conf.Env),
			attribute.Key("instance").String(p.Conf.Instance),
			attribute.Key("database_name").String(p.DatabaseName),
			attribute.Key("table").String(db.Statement.Table),
			attribute.Key("cmd").String(cmd),
		)
		MetricsMySQLRequestCounter.Add(context.TODO(), 1, opt)
	})
	_ = db.Callback().Row().After("gorm:row").Register("after:row", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		opt := api.WithAttributes(
			attribute.Key("app_id").String(p.Conf.AppID),
			attribute.Key("app_version").String(p.Conf.Version),
			attribute.Key("env").String(p.Conf.Env),
			attribute.Key("instance").String(p.Conf.Instance),
			attribute.Key("database_name").String(p.DatabaseName),
			attribute.Key("table").String(db.Statement.Table),
			attribute.Key("cmd").String(cmd),
		)
		MetricsMySQLRequestCounter.Add(context.TODO(), 1, opt)
	})

	_ = db.Callback().Raw().After("gorm:raw").Register("after:raw", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		opt := api.WithAttributes(
			attribute.Key("app_id").String(p.Conf.AppID),
			attribute.Key("app_version").String(p.Conf.Version),
			attribute.Key("env").String(p.Conf.Env),
			attribute.Key("instance").String(p.Conf.Instance),
			attribute.Key("database_name").String(p.DatabaseName),
			attribute.Key("table").String(db.Statement.Table),
			attribute.Key("cmd").String(cmd),
		)
		MetricsMySQLRequestCounter.Add(context.TODO(), 1, opt)
	})

	return nil
}
