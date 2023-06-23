package dao

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"web3/app/job/datawatch/storage/mq"
)

func (d *Dao) RabbitMQPushWeb3EthNewBlockMsg(ctx context.Context, data *mq.EthChainBlockHeader) error {
	blockByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = d.RabbitMQPublisherWeb3EthNewBlock.PublishWithContext(
		ctx,
		blockByte,
		[]string{uuid.NewString()},
		rabbitmq.WithPublishOptionsExchange("web3_new_block_eth"),
	)
	return err
}

func (d *Dao) RabbitMQPushWeb3EAptosNewBlockMsg(ctx context.Context, data *mq.AptosChainBlockInfo) error {
	blockByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = d.RabbitMQPublisherWeb3AptosNewBlock.PublishWithContext(
		ctx,
		blockByte,
		[]string{uuid.NewString()},
		rabbitmq.WithPublishOptionsExchange("web3_new_block_aptos"),
	)
	log.Context(ctx).Infow(
		"business", "rabbitmq",
		"exchange_name", "web3_aptos_event_block",
		"ms_data", string(blockByte),
		"error", err,
	)
	return err
}
