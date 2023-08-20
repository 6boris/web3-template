package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
	dexFinance "web3/api/service/dex-finance"
)

func TestService_Coin(t *testing.T) {
	t.Run("GetCoin", func(t *testing.T) {
		getResp, err := testSvc.GetCoin(testCtx, &dexFinance.GetCoinRequest{Id: 1012972373649002497})
		assert.Nil(t, err)
		spew.Dump(getResp)
	})
}
