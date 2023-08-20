package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "web3/api/service/dex-finance"
	"web3/app/service/dex-finance/internal/dao"
)

func (s *Service) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	resp := &pb.GetAccountReply{}
	accountInfo, err := dao.GlobalDao.BizGetAccountInfo(ctx, req.GetUserId(), req.GetCoinId(), req.GetAccountType())
	if err != nil {
		return nil, err
	}
	cTime := accountInfo.CreatedAt
	uTime := accountInfo.UpdatedAt
	resp.Item = &pb.AccountItem{
		Id:          accountInfo.ID,
		UserId:      accountInfo.UserID,
		CoinId:      accountInfo.CoinID,
		Status:      accountInfo.Status,
		AccountType: accountInfo.Type,
		Amount:      accountInfo.Amount,
		CreatedAt:   timestamppb.New(*cTime),
		UpdatedAt:   timestamppb.New(*uTime),
	}
	resp.ProtoReflect()
	return resp, nil
}

func (s *Service) AccountRechargeTry(ctx context.Context, req *pb.AccountRechargeRequest) (*pb.AccountRechargeReply, error) {
	resp := &pb.AccountRechargeReply{}
	_, err := dao.GlobalDao.BizGetAccountInfo(ctx, req.GetUserId(), req.GetCoinId(), req.GetAccountType())
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	accountInfo, err := dao.GlobalDao.BizRechargeAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	resp.UserId = accountInfo.UserID
	return resp, status.New(codes.FailedPrecondition, "").Err()
}
func (s *Service) AccountRechargeConfirm(ctx context.Context, req *pb.AccountRechargeRequest) (*pb.AccountRechargeReply, error) {
	resp := &pb.AccountRechargeReply{}
	_, err := dao.GlobalDao.BizGetAccountInfo(ctx, req.GetUserId(), req.GetCoinId(), req.GetAccountType())
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	accountInfo, err := dao.GlobalDao.BizRechargeAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	resp.UserId = accountInfo.UserID
	return resp, nil
}
func (s *Service) AccountRechargeCancel(ctx context.Context, req *pb.AccountRechargeRequest) (*pb.AccountRechargeReply, error) {
	resp := &pb.AccountRechargeReply{}
	_, err := dao.GlobalDao.BizGetAccountInfo(ctx, req.GetUserId(), req.GetCoinId(), req.GetAccountType())
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	accountInfo, err := dao.GlobalDao.BizRechargeAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	resp.UserId = accountInfo.UserID
	return resp, nil
}
