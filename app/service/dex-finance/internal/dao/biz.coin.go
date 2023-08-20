package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"web3/app/service/dex-finance/storage/mysql/model"
)

func (d *Dao) BizGetCoinInfo(ctx context.Context, id int64) (*model.Coin, error) {
	data := &model.Coin{
		ID: id,
	}
	var err error
	data, err = d.RedisGetCoinInfo(ctx, id)
	if err == nil {
		return data, nil
	}
	if err != redis.Nil {
		return nil, err
	}
	data, err = d.MySQLGetCoin(ctx, id)
	if err != nil {
		return nil, err
	}
	_ = d.RedisSetCoinInfo(ctx, data)
	return data, nil
}
