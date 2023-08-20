package client

import (
	"context"
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"testing"
	dexFinancePB "web3/api/service/dex-finance"
	pb "web3/api/service/dex-member"
	"web3/app/interface/dex/internal/conf"
	dexFinanceModel "web3/app/service/dex-finance/storage/mysql/model"
	"web3/pkg/app"
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
	app.GetAppLogger(app.APP_INTERFACE_DEX)
	testCtx = context.Background()
	testClient = New()
	//beforeTest(config.Conf)
	code := m.Run()
	//afterTest(config.Conf)
	os.Exit(code)
}

func Test_Unite(t *testing.T) {

	t.Run("GetUser", func(t *testing.T) {
		data, err := testClient.ClientAppServiceDexMemberUserGRPC.GetUser(context.Background(), &pb.GetUserRequest{Id: 105335491844780007})
		assert.Nil(t, err)
		spew.Dump(data)
	})
	t.Run("AuthSignIn", func(t *testing.T) {
		data, err := testClient.ClientAppServiceDexMemberAuthGRPC.AuthSignIn(testCtx, &pb.AuthSignInRequest{Email: "herryboy123@126.com"})
		assert.Nil(t, err)
		spew.Dump(data)
	})
	t.Run("RechargeAccount", func(t *testing.T) {
		req := &dexFinancePB.AccountRechargeRequest{
			UserId:      105335491844780007,
			CoinId:      100001,
			AccountType: dexFinanceModel.ACCOUNT_TYPE_BALANCE,
			TradeType:   dexFinanceModel.TRADE_TYPE_RECHARGE,
			Amount:      10.00001,
			StreamId:    uuid.NewString(),
			TradeTime:   timestamppb.Now(),
		}
		resp, err := testClient.ClientAppServiceDexFinanceAccountGRPC.AccountRechargeConfirm(testCtx, req)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
}
