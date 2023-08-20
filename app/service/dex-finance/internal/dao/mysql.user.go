package dao

import (
	"context"
	"web3/app/service/dex-finance/storage/mysql/model"
	"web3/app/service/dex-finance/storage/mysql/query"
)

func (d *Dao) MySQLGetCoin(ctx context.Context, id int64) (*model.Coin, error) {
	q := query.Use(d.MySQLClient).Coin
	return query.Use(d.MySQLClient).WithContext(ctx).Coin.
		Where(q.ID.Eq(id)).First()
}
