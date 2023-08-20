package dao

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"math/rand"
	"web3/app/service/dex-member/storage/mysql/model"
	"web3/app/service/dex-member/storage/mysql/query"
)

func (d *Dao) MySQLCreateUser(ctx context.Context, data *model.User) error {
	for i := 0; i < 10; i++ {
		data.ID = rand.Int63n(2223372036854775807)
		err := d.MySQLClient.WithContext(ctx).Create(data).Error
		if err == nil {
			return nil
		}
	}
	return errors.New(500, "SYSTEM_ERR", "generate uid fail")
}

func (d *Dao) MySQLGetUser(ctx context.Context, id int64) (*model.User, error) {
	q := query.Use(d.MySQLClient).User
	return query.Use(d.MySQLClient).WithContext(ctx).User.
		Where(q.ID.Eq(id)).First()
}

func (d *Dao) MySQLGetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	q := query.Use(d.MySQLClient).User
	return query.Use(d.MySQLClient).WithContext(ctx).User.
		Where(q.Email.Eq(email)).First()
}
