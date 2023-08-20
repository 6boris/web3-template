package client

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	dexFinance "web3/api/service/dex-finance"
	dexMember "web3/api/service/dex-member"
)

func Test_Dex_Unite(t *testing.T) {
	t.Run("CreateUser", func(t *testing.T) {
		createData := &dexMember.CreateUserRequest{
			Item: &dexMember.UserItem{
				NickName: uuid.NewString(),
			},
		}
		resp, err := testClient.ClientAppServiceDexMemberUserGRPC.CreateUser(testCtx, createData)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("GetCoin", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			getReq := &dexFinance.GetCoinRequest{
				Id: int64(i),
			}
			getResp, err := testClient.ClientAppServiceDexFinanceCoinGRPC.GetCoin(testCtx, getReq)
			assert.Nil(t, err)
			spew.Dump(getResp)
		}
	})

	t.Run("AuthSignIn", func(t *testing.T) {
		req := &dexMember.AuthSignInRequest{}
		resp, err := testClient.ClientAppServiceDexMemberAuthGRPC.AuthSignIn(testCtx, req)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
}
