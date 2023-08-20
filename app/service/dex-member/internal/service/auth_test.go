package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
	pb "web3/api/service/dex-member"
)

func TestService_Auth(t *testing.T) {
	t.Run("AuthSignIn", func(t *testing.T) {
		req := &pb.AuthSignInRequest{
			Email: "herryboy123@126.com",
		}
		resp, err := testSvc.AuthSignIn(testCtx, req)
		assert.Nil(t, err)
		spew.Dump(resp)

		verify, err := testSvc.AuthVerify(testCtx, &pb.AuthVerifyRequest{JwtToken: resp.GetJwtToken()})
		assert.Nil(t, err)
		spew.Dump(verify)
	})
}
