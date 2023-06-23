package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"web3/app/service/data/internal/client"
	"web3/app/service/data/storage/model"
)

func (d *Dao) BizDefiGetPrice(ctx context.Context, sourceToken, targetToken string) (*model.DefiPrice, error) {
	data := &model.DefiPrice{}
	var err error
	data, err = d.RedisGetDefiPrice(ctx, sourceToken, targetToken)
	if err == nil {
		return data, nil
	}
	if err != redis.Nil {
		return nil, err
	}
	data, err = client.GlobalClient.DefiGetTokenPrice(ctx, sourceToken, targetToken)
	if err != nil {
		return nil, err
	}
	_ = d.RedisSetDefiPrice(ctx, data)
	return data, err
}
