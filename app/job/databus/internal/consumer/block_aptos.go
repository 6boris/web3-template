package consumer

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wagslane/go-rabbitmq"
	"time"
	"web3/api/service/data"
	"web3/app/job/databus/internal/client"
	"web3/app/job/databus/internal/conf"
	"web3/app/job/datawatch/storage/mq"
)

type BlockAptos struct {
	RabbitMQConsume *rabbitmq.Consumer
}

func NewBlockAptos() *BlockAptos {
	return &BlockAptos{}
}

func (t *BlockAptos) Name() string {
	return "Consumer BlockAptos"
}
func (t *BlockAptos) Start(ctx context.Context) error {
	consumeConf := conf.Conf.GetRabbitmq().GetCWeb3NewBlockAptos()
	conn, err := rabbitmq.NewConn(
		consumeConf.GetUrl(),
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		panic(err)
	}
	consumer, err := rabbitmq.NewConsumer(
		conn,
		t._syncBlockData,
		consumeConf.GetQueueName(),
		rabbitmq.WithConsumerOptionsRoutingKey(""),
		rabbitmq.WithConsumerOptionsExchangeName(consumeConf.GetExchangeName()),
		rabbitmq.WithConsumerOptionsExchangeKind(consumeConf.GetExchangeKind()),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsConcurrency(int(consumeConf.GetConcurrency())),
	)
	if err != nil {
		log.Context(ctx).Errorf("[Consume] BlockAptos err:%+v", err)
	}
	t.RabbitMQConsume = consumer
	return nil
}
func (t *BlockAptos) Stop(ctx context.Context) error {
	t.RabbitMQConsume.Close()
	return nil
}
func (t *BlockAptos) Reload(ctx context.Context) error {
	return nil
}
func (t *BlockAptos) _syncBlockData(d rabbitmq.Delivery) rabbitmq.Action {
	log.Info("consumed: %+v", string(d.Body))
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	var err error
	msgData := &mq.AptosChainBlockInfo{}
	err = json.Unmarshal(d.Body, msgData)
	if err != nil {
		log.Errorf("err: %v", err)
		return rabbitmq.Ack
	}
	saveReq := &data.SyncBlockInfoRequest{Item: &data.BlockItem{
		Number:            msgData.BlockHeight,
		Hash:              msgData.BlockHash,
		ParentHash:        common.HexToHash("").String(),
		Timestamp:         msgData.BlockTimestamp / 1e6,
		Difficulty:        "",
		ExtraData:         "",
		GasLimit:          "",
		GasUsed:           "",
		BaseFeePerGas:     "",
		Miner:             common.HexToAddress("").Hex(),
		MixHash:           common.HexToHash("").Hex(),
		Nonce:             "",
		ReceiptsRoot:      common.HexToHash("").Hex(),
		Sha3Uncles:        common.HexToHash("").Hex(),
		StateRoot:         common.HexToHash("").Hex(),
		Size:              0,
		TransactionsRoot:  common.HexToHash("").Hex(),
		TransactionsCount: int64(len(msgData.Transactions)),
		UnclesCount:       0,
		ChainId:           msgData.ChainID,
		ChainUniqueId:     int64(msgData.ChainUniqueID),
	}}
	saveResp, err := client.GlobalClient.ClientAppServiceDataChainRPC.SyncBlockInfo(ctx, saveReq)
	if err != nil {
		log.Errorf("SyncBlockInfo saveReq: %+v saveResp: %+v err: %+v err: %+v", saveReq, saveResp, err)
		return rabbitmq.Ack
	}
	for idx, tx := range msgData.Transactions {
		saveTxReq := &data.SyncBlockTransactionInfoRequest{
			Item: &data.TransactionItem{
				Type:                 0,
				Status:               1,
				BlockHash:            msgData.BlockHash,
				BlockNumber:          msgData.BlockHeight,
				BlockTimestamp:       msgData.BlockTimestamp / 1e6,
				TransactionHash:      tx.Hash,
				TransactionIndex:     int64(idx),
				FromAddress:          common.HexToAddress("").String(),
				ToAddress:            common.HexToAddress("").String(),
				Value:                "",
				Input:                "",
				Nonce:                0,
				ContractAddress:      common.HexToAddress("").String(),
				Gas:                  0,
				GasPrice:             0,
				GasUsed:              0,
				EffectiveGasPrice:    0,
				CumulativeGasUsed:    0,
				MaxFeePerGas:         0,
				MaxPriorityFeePerGas: 0,
				R:                    "",
				S:                    "",
				V:                    0,
				LogsCount:            0,
				ChainId:              msgData.ChainID,
				ChainUniqueId:        int64(msgData.ChainUniqueID),
			},
		}

		saveTxResp, loopErr := client.GlobalClient.ClientAppServiceDataChainRPC.SyncBlockTransactionInfo(ctx, saveTxReq)
		if err != nil {
			log.Errorf("SaveTransaction saveReq: %+v saveResp: %+v err: %+v err: %+v", saveReq, saveTxResp, loopErr)
			continue
		}
	}

	return rabbitmq.Ack
}
