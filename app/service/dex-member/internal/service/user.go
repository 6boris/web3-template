package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "web3/api/service/dex-member"
	"web3/app/service/dex-member/internal/dao"
	"web3/app/service/dex-member/storage/mysql/model"
)

func (s *Service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return nil, errors.New(500, "SYSTEM_ERR", "not support this opt")
}

func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return nil, errors.New(500, "SYSTEM_ERR", "not support this opt")
}
func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	createData := &model.User{
		NickName:  req.GetItem().GetNickName(),
		AvatarURL: req.GetItem().GetAvatarUrl(),
		Email:     req.GetItem().GetEmail(),
	}
	err := dao.GlobalDao.MySQLCreateUser(ctx, createData)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{Id: createData.ID}, nil
}
func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	resp := &pb.GetUserReply{
		Item: &pb.UserItem{},
	}
	user, err := s.dao.BizGetAccount(ctx, req.GetId())
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	cTime := user.CreatedAt
	uTime := user.UpdatedAt
	resp.Item.Id = user.ID
	resp.Item.NickName = user.NickName
	resp.Item.AvatarUrl = user.AvatarURL
	resp.Item.Email = user.Email
	resp.Item.CreatedAt = timestamppb.New(*cTime)
	resp.Item.UpdatedAt = timestamppb.New(*uTime)
	return resp, nil
}

func (s *Service) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return nil, errors.New(500, "SYSTEM_ERR", "not support this opt")
}
