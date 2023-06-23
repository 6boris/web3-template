package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"web3/app/service/data/storage/model"
	"web3/app/service/data/storage/redis"
)

func (d *Dao) RedisSetDefiPrice(ctx context.Context, data *model.DefiPrice) error {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("web3:data:defi_price:%s:%s", data.SourceToken, data.TargetToken),
		TTL: 60,
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return d.RedisClient.SetEx(ctx, cmd.Key, string(dataByte), time.Second*time.Duration(cmd.TTL)).Err()
}

func (d *Dao) RedisGetDefiPrice(ctx context.Context, sourceToken, targetToken string) (*model.DefiPrice, error) {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("web3:data:defi_price:%s:%s", sourceToken, targetToken),
		TTL: 3600,
	}
	data := &model.DefiPrice{}
	dataBytes, err := d.RedisClient.GetEx(ctx, cmd.Key, time.Second*time.Duration(cmd.TTL)).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, data)
	return data, err
}
