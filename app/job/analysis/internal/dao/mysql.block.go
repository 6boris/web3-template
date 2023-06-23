package dao

import (
	"context"
	"time"
	"web3/app/job/analysis/storage/mysql/query"
)

func (d *Dao) MySQLDeleteBlockHistoryData(ctx context.Context, deadline time.Time) (int64, error) {
	var err error
	q := query.Use(d.Web3MySQLClient).Block
	data, err := q.WithContext(ctx).Where(q.Timestamp.Lt(deadline)).Delete()
	if err != nil {
		return 0, err
	}
	return data.RowsAffected, nil
}

func (d *Dao) MySQLDeleteBlockTransactionHistoryData(ctx context.Context, deadline time.Time) (int64, error) {
	var err error
	q := query.Use(d.Web3MySQLClient).BlockTransaction
	data, err := q.WithContext(ctx).Where(q.BlockTimestamp.Lt(deadline)).Delete()
	if err != nil {
		return 0, err
	}
	return data.RowsAffected, nil
}
