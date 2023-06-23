package dao

import (
	"context"
	"web3/app/service/data/storage/mysql/model"
	"web3/app/service/demo/storage/mysql/query"
)

func (d *Dao) MySQLSyncBlock(ctx context.Context, data *model.Block) error {
	q := query.Use(d.MySQLClient).Block
	return d.MySQLClient.WithContext(ctx).Where(
		q.ChainUniqueID.Eq(data.ChainUniqueID),
		q.Number.Eq(data.Number),
	).FirstOrCreate(data).Error
}

func (d *Dao) MySQLSyncBlockTransaction(ctx context.Context, data *model.BlockTransaction) error {
	q := query.Use(d.MySQLClient).BlockTransaction
	return d.MySQLClient.WithContext(ctx).Where(
		q.ChainUniqueID.Eq(data.ChainUniqueID),
		q.TransactionHash.Eq(data.TransactionHash),
	).FirstOrCreate(data).Error
}
