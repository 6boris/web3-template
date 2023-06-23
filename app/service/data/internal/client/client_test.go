package client

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"os"
	"testing"
	"web3/app/service/data/internal/conf"
)

var (
	testClient *Client
	testCtx    context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	if err := os.Setenv("ENV", "dev"); err != nil {
		return
	}
	conf.Init("../../configs")
	bc := conf.Conf
	logger := log.With(
		log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"source", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
		"service.id", bc.GetApp().GetId(),
		"service.env", bc.GetApp().GetEnv(),
		"service.instance", "",
		"service.cluster", bc.GetApp().GetCluster(),
		"service.zone", bc.GetApp().GetZone(),
		"service.version", bc.GetApp().GetVersion(),
	)
	testClient = New(conf.Conf, logger)
	testCtx = context.Background()
	//beforeTest(config.Conf)
	code := m.Run()
	//afterTest(config.Conf)
	os.Exit(code)
}
