package dao

import (
	"context"
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"web3/app/service/demo/internal/conf"
)

var (
	testDao *Dao
	testCtx context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	if err := os.Setenv("ENV", "dev"); err != nil {
		return
	}
	conf.Init("../../configs")
	testDao = New(conf.Conf)
	testCtx = context.Background()
	//beforeTest(config.Conf)
	code := m.Run()
	//afterTest(config.Conf)

	os.Exit(code)
}

func TestPing(t *testing.T) {
	var err error
	err = testDao.RedisClient.Ping(testCtx).Err()
	assert.Nil(t, err)
}
