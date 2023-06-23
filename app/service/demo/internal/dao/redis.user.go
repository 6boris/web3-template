package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"web3/app/service/demo/storage/mysql/model"
	"web3/app/service/demo/storage/redis"
)

func (d *Dao) RedisSetUserInfo(ctx context.Context, data *model.User) error {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("web3:demo:user_info:%d", data.ID),
		TTL: 86400,
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return d.RedisClient.SetEx(ctx, cmd.Key, string(dataByte), time.Second*time.Duration(cmd.TTL)).Err()
}

func (d *Dao) RedisGetUserInfo(ctx context.Context, id int64) (*model.User, error) {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("web3:demo:user_info:%d", id),
		TTL: 86400,
	}
	data := &model.User{ID: id}
	dataBytes, err := d.RedisClient.GetEx(ctx, cmd.Key, time.Second*time.Duration(cmd.TTL)).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, data)
	return data, err
}

func (d *Dao) RedisDelUserInfo(ctx context.Context, id int64) error {
	cmd := redis.Cmd{
		Key: fmt.Sprintf("web3:demo:user_info:%d", id),
		TTL: 86400,
	}
	err := d.RedisClient.Del(ctx, cmd.Key).Err()
	if err != nil {
		return err
	}
	return nil
}
