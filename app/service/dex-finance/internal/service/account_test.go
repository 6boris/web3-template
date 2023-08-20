package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"testing"
	dexFinance "web3/api/service/dex-finance"
	"web3/app/service/dex-finance/storage/mysql/model"
)

func TestService_Account(t *testing.T) {
	accountTypes := []string{
		model.ACCOUNT_TYPE_DEFAULT,
		model.ACCOUNT_TYPE_BALANCE,
		model.ACCOUNT_TYPE_NET_VALUE,
		model.ACCOUNT_TYPE_FREEZE,
		model.ACCOUNT_TYPE_TOTAL_RECHARGE,
		model.ACCOUNT_TYPE_TOTAL_WITHDRAW,
		model.ACCOUNT_TYPE_FLOAT_PROFIT,
		model.ACCOUNT_TYPE_TOTAL_PROFIT,
	}
	tradeTypes := []string{
		model.TRADE_TYPE_DEFAULT,
		model.TRADE_TYPE_RECHARGE,
		model.TRADE_TYPE_WITHDRAWAL,
		model.TRADE_TYPE_TRANS_IN,
		model.TRADE_TYPE_TRANS_OUT,
	}
	t.Run("GetAccount", func(t *testing.T) {
		for _, v := range accountTypes {
			getResp, err := testSvc.GetAccount(testCtx, &dexFinance.GetAccountRequest{
				UserId:      20154913779637377,
				CoinId:      1001784310992068611,
				AccountType: v,
			})
			assert.Nil(t, err)
			spew.Dump(getResp)
		}
	})
	t.Run("RechargeAccount", func(t *testing.T) {
		for _, v := range accountTypes {
			getResp, err := testSvc.RechargeAccount(testCtx, &dexFinance.RechargeAccountRequest{
				UserId:      20154913779637377,
				CoinId:      1001784310992068612,
				AccountType: v,
				Amount:      rand.Float64(),
				StreamId:    uuid.NewString(),
			})
			assert.Nil(t, err)
			spew.Dump(getResp)
		}

	})
	t.Run("RechargeAccount", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			for _, at := range accountTypes {
				for _, tt := range tradeTypes {
					getResp, err := testSvc.RechargeAccount(testCtx, &dexFinance.RechargeAccountRequest{
						UserId:      rand.Int63n(20154913779637377),
						CoinId:      rand.Int63n(20154913779637377),
						AccountType: at,
						TradeType:   tt,
						Amount:      rand.Float64() - rand.Float64(),
						StreamId:    uuid.NewString(),
						Remark:      "remark",
						ExtraData:   "extra_data",
						TradeTime:   timestamppb.Now(),
					})
					assert.Nil(t, err)
					spew.Dump(getResp)
				}
			}
		}
	})
}
