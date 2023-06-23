package demo

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	pb "web3/api/interface/web3"
	"web3/app/interface/web3/internal/client"
)

func (s *DemoService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	resp := &pb.GetUserReply{
		Item: &pb.GetUserReply_Item{},
	}
	userItem, err := client.GlobalClient.DemoGetUser(ctx, req.GetId())
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	resp.Item.Id = userItem.GetId()
	resp.Item.NickName = userItem.GetNickName()
	resp.Item.AvatarUrl = userItem.GetAvatarUrl()
	resp.Item.Email = userItem.GetEmail()
	resp.Item.CreatedAt = userItem.GetCreatedAt()
	resp.Item.UpdatedAt = userItem.GetUpdatedAt()
	return resp, nil
}
func (s *DemoService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	resp := &pb.CreateUserReply{}
	userID, err := client.GlobalClient.DemoCreateUser(ctx)
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	resp.Id = userID
	return resp, nil
}
