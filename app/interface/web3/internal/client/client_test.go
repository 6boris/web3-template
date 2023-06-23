package client

import (
	"context"
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	pb "web3/api/service/demo"
	"web3/app/interface/web3/internal/conf"
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

func TestNew(t *testing.T) {
	t.Run("UserGRPC", func(t *testing.T) {
		data, err := testClient.ClientAppServiceDemoUserGRPC.GetUser(testCtx, &pb.GetUserRequest{Id: 1})
		assert.Nil(t, err)
		spew.Dump(data)
	})
	t.Run("UserHTTP", func(t *testing.T) {
		data, err := testClient.ClientAppServiceDemoUserHTTP.GetUser(testCtx, &pb.GetUserRequest{Id: 1})
		assert.Nil(t, err)
		spew.Dump(data)
	})
}
