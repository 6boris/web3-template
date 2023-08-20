package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"web3/app/service/dex-finance/storage/mysql/model"
	"web3/app/service/dex-finance/storage/redis"
)

func (d *Dao) RedisSetCoinInfo(ctx context.Context, data *model.Coin) error {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("dex:finance:coin_info:%d", data.ID),
		TTL: 86400,
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return d.RedisClient.SetEx(ctx, cmd.Key, string(dataByte), time.Second*time.Duration(cmd.TTL)).Err()
}

func (d *Dao) RedisGetCoinInfo(ctx context.Context, id int64) (*model.Coin, error) {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("dex:finance:coin_info:%d", id),
		TTL: 86400,
	}
	data := &model.Coin{ID: id}
	dataBytes, err := d.RedisClient.GetEx(ctx, cmd.Key, time.Second*time.Duration(cmd.TTL)).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, data)
	return data, err
}

func (d *Dao) RedisDelCoinInfo(ctx context.Context, id int64) error {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("dex:finance:coin_info:%d", id),
		TTL: 86400,
	}
	err := d.RedisClient.Del(ctx, cmd.Key).Err()
	if err != nil {
		return err
	}
	return nil
}
