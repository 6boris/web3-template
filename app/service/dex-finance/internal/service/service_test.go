package service

import (
	"context"
	"flag"
	"os"
	"testing"
	"web3/app/service/dex-finance/internal/conf"
	"web3/app/service/dex-finance/internal/dao"
)

var (
	testSvc *Service
	testCtx context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	if err := os.Setenv("ENV", "dev"); err != nil {
		return
	}
	conf.Init("../../configs")
	testSvc = New(conf.Conf)
	dao.New(conf.Conf)
	testCtx = context.Background()
	//beforeTest(config.Conf)
	code := m.Run()
	//afterTest(config.Conf)

	os.Exit(code)
}
