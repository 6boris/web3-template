package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
	pb "web3/api/service/dex-finance"
	"web3/app/service/dex-finance/storage/mysql/model"
	"web3/app/service/dex-finance/storage/mysql/query"
)

func (d *Dao) MySQLGetOrCreateAccount(ctx context.Context, userID, coinID int64, accountType string) (*model.Account, error) {
	q := query.Use(d.MySQLClient).Account
	return query.Use(d.MySQLClient).WithContext(ctx).Account.
		Assign(
			q.UserID.Value(userID),
			q.CoinID.Value(coinID),
			q.Type.Value(accountType),
			q.Status.Value(1),
			q.Amount.Value(0),
			q.CreatedAt.Value(time.Now()),
			q.UpdatedAt.Value(time.Now()),
		).
		Where(
			q.UserID.Eq(userID),
			q.CoinID.Eq(coinID),
			q.Type.Eq(accountType),
		).FirstOrCreate()
}

func (d *Dao) MySQLRechargeAccount(ctx context.Context, rechargeReq *pb.AccountRechargeRequest) error {
	userID := rechargeReq.GetUserId()
	coinID := rechargeReq.GetCoinId()
	accountType := rechargeReq.GetAccountType()
	tradeType := rechargeReq.GetTradeType()
	amount := rechargeReq.GetAmount()
	streamID := rechargeReq.GetStreamId()
	direction := int64(1)
	remark := rechargeReq.GetRemark()
	extraData := rechargeReq.GetExtraData()
	tradeTime := rechargeReq.GetTradeTime().AsTime()

	q := query.Use(d.MySQLClient)
	tmpErr := q.Transaction(func(tx *query.Query) error {
		qat := query.Use(d.MySQLClient).AccountTrade
		_, err := tx.AccountTrade.WithContext(ctx).Where(qat.StreamID.Eq(streamID)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if err != gorm.ErrRecordNotFound {
			return nil
		}

		accountInfo, err := tx.Account.WithContext(ctx).Where(
			q.Account.UserID.Eq(userID),
			q.Account.CoinID.Eq(coinID),
			q.Account.Type.Eq(accountType),
		).First()
		if err != nil {
			return err
		}
		if amount < 0 {
			direction = 2
		}
		_, err = tx.Account.WithContext(ctx).Where(
			q.Account.UserID.Eq(userID),
			q.Account.CoinID.Eq(coinID),
			q.Account.Type.Eq(accountType),
		).UpdateSimple(
			q.Account.Amount.Add(amount),
		)
		if err != nil {
			return err
		}
		accountTrade := &model.AccountTrade{
			UserID:      userID,
			CoinID:      coinID,
			AccountID:   accountInfo.ID,
			StreamID:    streamID,
			AccountType: accountInfo.Type,
			TradeType:   tradeType,
			Direction:   direction,
			Amount:      amount,
			Remark:      remark,
			Status:      1,
			ExtraData:   extraData,
			TradeTime:   &tradeTime,
		}
		err = tx.AccountTrade.WithContext(ctx).Create(accountTrade)
		if err != nil {
			return err
		}
		return nil
	})
	return tmpErr
}
