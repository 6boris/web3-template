package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "web3/api/service/dex-finance"
	"web3/app/service/dex-finance/internal/dao"
)

func (s *Service) GetCoin(ctx context.Context, req *pb.GetCoinRequest) (*pb.GetCoinReply, error) {
	resp := &pb.GetCoinReply{
		Item: &pb.CoinItem{},
	}
	user, err := dao.GlobalDao.BizGetCoinInfo(ctx, req.GetId())
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	cTime := user.CreatedAt
	uTime := user.UpdatedAt
	resp.Item.Id = user.ID
	resp.Item.Name = user.Name
	resp.Item.Title = user.Title
	resp.Item.Img = user.Img
	resp.Item.Round = user.Round
	resp.Item.Title = user.Title
	resp.Item.BaseAmount = *user.BaseAmount
	resp.Item.CreatedAt = timestamppb.New(*cTime)
	resp.Item.UpdatedAt = timestamppb.New(*uTime)
	return resp, nil
}
