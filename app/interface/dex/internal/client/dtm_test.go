package client

import (
	"fmt"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/dtm-labs/dtm/dtmutil"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
	dexFinancePB "web3/api/service/dex-finance"
	dexFinanceModel "web3/app/service/dex-finance/storage/mysql/model"
)

func Test_DTM(t *testing.T) {
	//var DtmServer = "http://localhost:36789/api/dtmsvr"
	var BusiGrpc = "127.0.0.1:61001"
	t.Run("TCC", func(t *testing.T) {
		gid := dtmgrpc.MustGenGid(dtmutil.DefaultGrpcServer)
		fmt.Println(gid)
		err := dtmgrpc.TccGlobalTransaction(dtmutil.DefaultGrpcServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
			reqData := &dexFinancePB.AccountRechargeRequest{
				UserId:      105335491844780007,
				CoinId:      100001,
				AccountType: dexFinanceModel.ACCOUNT_TYPE_BALANCE,
				TradeType:   dexFinanceModel.TRADE_TYPE_RECHARGE,
				Amount:      10.00001,
				StreamId:    dtmgrpc.MustGenGid(dtmutil.DefaultGrpcServer),
				TradeTime:   timestamppb.Now(),
				Remark:      "",
				ExtraData:   "",
			}
			respData := &dexFinancePB.AccountRechargeReply{}
			err := tcc.CallBranch(
				reqData,
				BusiGrpc+"/api.service.dex_finance.Account/AccountRechargeTry",
				BusiGrpc+"/api.service.dex_finance.Account/AccountRechargeConfirm",
				BusiGrpc+"/api.service.dex_finance.Account/AccountRechargeCancel",
				respData,
			)
			err = tcc.CallBranch(
				reqData,
				BusiGrpc+"/api.service.dex_finance.Account/AccountRechargeTry",
				BusiGrpc+"/api.service.dex_finance.Account/AccountRechargeConfirm",
				BusiGrpc+"/api.service.dex_finance.Account/AccountRechargeCancel",
				respData,
			)
			return err
		})
		assert.Nil(t, err)
		time.Sleep(time.Second)
	})
}
