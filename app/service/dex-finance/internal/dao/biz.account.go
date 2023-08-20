package dao

import (
	"context"
	pb "web3/api/service/dex-finance"
	"web3/app/service/dex-finance/storage/mysql/model"
)

func (d *Dao) BizGetAccountInfo(ctx context.Context, userID, coinID int64, accountType string) (*model.Account, error) {
	data := &model.Account{}
	data, err := d.MySQLGetOrCreateAccount(ctx, userID, coinID, accountType)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *Dao) BizRechargeAccount(ctx context.Context, rechargeReq *pb.AccountRechargeRequest) (*model.Account, error) {
	data := &model.Account{}
	err := d.MySQLRechargeAccount(ctx, rechargeReq)
	if err != nil {
		return nil, err
	}
	return data, nil
}
