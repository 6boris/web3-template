package dao

import (
	"context"
	"github.com/wagslane/go-rabbitmq"
	"web3/app/job/datawatch/internal/conf"
)

var GlobalDao *Dao

// Dao ...
type Dao struct {
	RabbitMQPublisherWeb3EthNewBlock   *rabbitmq.Publisher
	RabbitMQPublisherWeb3AptosNewBlock *rabbitmq.Publisher
}

// New ...
func New(conf *conf.Bootstrap) *Dao {
	d := &Dao{}
	var err error
	blockPublisherConf := conf.GetRabbitmq().GetPWeb3NewBlockEth()
	conn1, err := rabbitmq.NewConn(
		blockPublisherConf.GetUrl(),
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		panic(err)
	}
	d.RabbitMQPublisherWeb3EthNewBlock, err = rabbitmq.NewPublisher(
		conn1,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(blockPublisherConf.GetExchangeName()),
		rabbitmq.WithPublisherOptionsExchangeKind(blockPublisherConf.GetExchangeKind()),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		panic(err)
	}
	conn2, err := rabbitmq.NewConn(
		blockPublisherConf.GetUrl(),
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		panic(err)
	}
	aptBlockPublisherConf := conf.GetRabbitmq().GetPWeb3NewBlockAptos()
	d.RabbitMQPublisherWeb3AptosNewBlock, err = rabbitmq.NewPublisher(
		conn2,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(aptBlockPublisherConf.GetExchangeName()),
		rabbitmq.WithPublisherOptionsExchangeKind(aptBlockPublisherConf.GetExchangeKind()),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		panic(err)
	}
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
