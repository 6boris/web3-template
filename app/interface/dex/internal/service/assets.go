package service

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/rand"
	pb "web3/api/interface/dex"
	dexMember "web3/api/service/dex-member"
	"web3/app/interface/dex/internal/client"
)

func (s *Service) GetUserAccountsAssets(ctx context.Context, req *pb.GetUserAccountsAssetsRequest) (*pb.GetUserAccountsAssetsReply, error) {
	resp := &pb.GetUserAccountsAssetsReply{}

	verifyReq := &dexMember.AuthVerifyRequest{JwtToken: req.GetJwtToken()}
	verifyResp, err := client.GlobalClient.ClientAppServiceDexMemberAuthGRPC.AuthVerify(ctx, verifyReq)
	if err != nil {
		return nil, err
	}
	resp.User = &pb.GetUserAccountsAssetsReply_User{
		UserId:   verifyResp.GetUser().GetId(),
		NickName: verifyResp.GetUser().GetNickName(),
	}
	resp.Accounts = []*pb.GetUserAccountsAssetsReply_Account{}
	for i := 1; i <= 10; i++ {
		resp.Accounts = append(resp.Accounts, &pb.GetUserAccountsAssetsReply_Account{
			Rank:            int64(i),
			UserId:          resp.User.UserId,
			CoinId:          rand.Int63n(100000),
			Type:            1,
			Status:          1,
			BalanceAmount:   rand.Float64(),
			FreezeAmount:    rand.Float64(),
			RechargeAddress: common.HexToAddress("").String(),
		})
	}

	return resp, nil
}
