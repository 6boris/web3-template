package dao

import (
	"context"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"time"
	"web3/app/job/analysis/storage/mysql/model"
	"web3/app/job/analysis/storage/mysql/query"
)

func (d *Dao) MySQLSaveChainMetric(ctx context.Context, data *model.ChainMetric) error {
	var err error
	q := query.Use(d.Web3MySQLClient).ChainMetric
	_, err = q.WithContext(ctx).Where(q.ChainUniqueID.Eq(data.ChainUniqueID)).First()
	if err == gorm.ErrRecordNotFound {
		return q.WithContext(ctx).Create(data)
	}
	updateFileds := make([]field.AssignExpr, 0)
	if data.GasPrice != 0 {
		updateFileds = append(updateFileds, q.GasPrice.Value(data.GasPrice))
	}
	if data.LatestBlockNumber != 0 {
		updateFileds = append(updateFileds, q.LatestBlockNumber.Value(data.LatestBlockNumber))
	}
	if data.TxTps != 0 {
		updateFileds = append(updateFileds, q.TxTps.Value(data.TxTps))
	}
	if data.TxTpd != 0 {
		updateFileds = append(updateFileds, q.TxTpd.Value(data.TxTpd))
	}
	if data.TxPending != 0 {
		updateFileds = append(updateFileds, q.TxPending.Value(data.TxPending))
	}
	if data.TokenPriceUsd != 0 {
		updateFileds = append(updateFileds, q.TokenPriceUsd.Value(data.TokenPriceUsd))
	}
	_, err = q.WithContext(ctx).Where(q.ChainUniqueID.Eq(data.ChainUniqueID)).
		UpdateColumnSimple(updateFileds...)
	return err
}

func (d *Dao) MySQLGetRecentTxMetrics(ctx context.Context, chainUniqueID int64, startAt time.Time, endAt time.Time, dur int64) (float64, error) {
	type TxTps struct {
		Id  int64   `gorm:"column:id;"`
		Tps float64 `gorm:"column:tpx;"`
	}
	data := &TxTps{}
	q := query.Use(d.Web3MySQLClient).Block
	err := q.WithContext(ctx).
		Where(
			q.ChainUniqueID.Eq(chainUniqueID),
			q.Timestamp.Gte(startAt),
			q.Timestamp.Lte(endAt),
		).
		Select(
			q.ChainUniqueID.As("id"),
			q.TransactionsCount.Sum().Div(dur).As("tpx")).
		Group(q.ChainUniqueID).Scan(data)
	return data.Tps, err
}
