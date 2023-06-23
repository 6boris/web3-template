package task

import (
	"context"
	"flag"
	"os"
	"testing"
	"web3/app/job/datawatch/internal/conf"
)

var (
	testSrv *Server
	testCtx context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	if err := os.Setenv("ENV", "dev"); err != nil {
		return
	}
	conf.Init("../../configs")
	testSrv = NewServer()
	testCtx = context.Background()
	//beforeTest(config.Conf)
	code := m.Run()
	//afterTest(config.Conf)

	os.Exit(code)
}
