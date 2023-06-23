package dao

import (
	"context"
	"web3/app/job/databus/internal/conf"
)

var GlobalDao *Dao

// Dao ...
type Dao struct {
}

// New ...
func New(conf *conf.Bootstrap) *Dao {
	d := &Dao{}
	GlobalDao = d
	return d
}

func (d *Dao) Close(ctx context.Context) error {
	var err error
	return err
}
func (d *Dao) Ping(ctx context.Context) error {
	var err error
	return err
}
