package dao

import (
	"context"
	"web3/app/job/analysis/storage/mysql/model"
	"web3/app/job/analysis/storage/mysql/query"
)

func (d *Dao) MySQLSearchChainInfoList(ctx context.Context, uniqueIds []int64) ([]*model.ChainInfo, error) {
	var err error
	q := query.Use(d.Web3MySQLClient).ChainInfo
	data, err := q.WithContext(ctx).Where(q.UniqueID.In(uniqueIds...)).
		Order(q.OrderSeq.Desc()).Limit(100).Find()
	return data, err
}

func (d *Dao) MySQLGetAllChainInfoList(ctx context.Context) ([]*model.ChainInfo, error) {
	var err error
	q := query.Use(d.Web3MySQLClient).ChainInfo
	data, err := q.WithContext(ctx).Find()
	return data, err
}
