package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"web3/app/service/dex-member/storage/mysql/model"
)

func (d *Dao) BizGetAccount(ctx context.Context, id int64) (*model.User, error) {
	data := &model.User{
		ID: id,
	}
	var err error
	data, err = d.RedisGetUserInfo(ctx, id)
	if err == nil {
		return data, nil
	}
	if err != redis.Nil {
		return nil, err
	}
	data, err = d.MySQLGetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	_ = d.RedisSetUserInfo(ctx, data)
	return data, nil
}

func (d *Dao) BizCreateUser(ctx context.Context, data *model.User) error {
	return d.MySQLCreateUser(ctx, data)
}
