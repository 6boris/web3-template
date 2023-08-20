package app

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	transportGrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transportHttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"os"
	"strings"
	"web3/pkg/otel"
)

type APP_SERVICE_NAME string

const (
	APP_SERVICE_DEMO        APP_SERVICE_NAME = "app.service.demo"
	APP_SERVICE_DATA        APP_SERVICE_NAME = "app.service.data"
	APP_SERVICE_DEX_FINANCE APP_SERVICE_NAME = "app.service.dex-finance"
	APP_SERVICE_DEX_MEMBER  APP_SERVICE_NAME = "app.service.dex-member"
	APP_INTERFACE_DEX       APP_SERVICE_NAME = "app.interface.dex"
	APP_INTERFACE_WEB3      APP_SERVICE_NAME = "app.interface.web3"
	APP_JOB_ANALYSIS        APP_SERVICE_NAME = "app.job.analysis"
	APP_JOB_DATABUS         APP_SERVICE_NAME = "app.job.databus"
	APP_JOB_DATAWATCH       APP_SERVICE_NAME = "app.job.datawatch"

	_defaultAppServiceEnv     = "default.env"
	_defaultAppServiceRegion  = "default.region"
	_defaultAppServiceZone    = "default.zone"
	_defaultAppServiceCluster = "default.cluster"
	_defaultAppServiceVersion = "default.version"
)

var _defaultHostName, _ = os.Hostname()
var _defaultAppServiceId = fmt.Sprintf("%s:%s", _defaultHostName, strings.ReplaceAll(uuid.NewString(), "-", ""))
var _defaultEtcdOpts = []etcd.Option{etcd.Namespace("/web3-template/microservices")}

func (a APP_SERVICE_NAME) String() string {
	return string(a)
}

func GetAppOpts(name APP_SERVICE_NAME) []kratos.Option {
	logger := GetAppLogger(name)
	opts := []kratos.Option{
		kratos.ID(_defaultAppServiceId),
		kratos.Name(string(name)),
		kratos.Version("default.version"),
		kratos.Metadata(map[string]string{
			"hostname": _defaultAppServiceId,
			"appid":    string(name),
			"env":      _defaultAppServiceEnv,
			"region":   _defaultAppServiceRegion,
			"zone":     _defaultAppServiceZone,
			"cluster":  _defaultAppServiceCluster,
			"version":  _defaultAppServiceVersion,
		}),
		kratos.Logger(logger),
		kratos.Registrar(etcd.New(GetRegistryEtcdClientWithDefalut(), _defaultEtcdOpts...)),
	}
	return opts
}

func GetAppLogger(name APP_SERVICE_NAME) log.Logger {
	logger := log.With(
		log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"source", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
		"service.id", _defaultAppServiceId,
		"service.name", string(name),
		"service.env", _defaultAppServiceEnv,
		"service.region", _defaultAppServiceRegion,
		"service.zone", _defaultAppServiceZone,
		"service.cluster", _defaultAppServiceCluster,
		"service.version", _defaultAppServiceVersion,
	)
	return logger
}

func GetGrpcClientConn(name APP_SERVICE_NAME) (*grpc.ClientConn, error) {
	logger := GetAppLogger(name)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	conn, err := transportGrpc.DialInsecure(context.Background(),
		transportGrpc.WithEndpoint(
			fmt.Sprintf("discovery:///%s", name.String()),
		),
		transportGrpc.WithDiscovery(etcd.New(cli, _defaultEtcdOpts...)),
		transportGrpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			logging.Client(logger),
			otel.MetricsClient(&otel.Conf{
				AppID:    name.String(),
				Env:      _defaultAppServiceEnv,
				Instance: _defaultAppServiceId,
				Cluster:  _defaultAppServiceCluster,
				Zone:     _defaultAppServiceZone,
				Version:  _defaultAppServiceVersion,
			}),
		),
	)
	return conn, err
}
func GetHttpClientConn(name APP_SERVICE_NAME) (*transportHttp.Client, error) {
	logger := GetAppLogger(name)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	conn, err := transportHttp.NewClient(context.Background(),
		transportHttp.WithEndpoint(
			fmt.Sprintf("discovery:///%s", name.String()),
		),
		transportHttp.WithDiscovery(etcd.New(cli, _defaultEtcdOpts...)),
		transportHttp.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			logging.Client(logger),
			otel.MetricsClient(&otel.Conf{
				AppID:    name.String(),
				Env:      _defaultAppServiceEnv,
				Instance: _defaultAppServiceId,
				Cluster:  _defaultAppServiceCluster,
				Zone:     _defaultAppServiceZone,
				Version:  _defaultAppServiceVersion,
			}),
		),
	)
	return conn, err
}
