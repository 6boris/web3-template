package consumer

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wagslane/go-rabbitmq"
	"math/big"
	"strconv"
	"web3/api/service/data"
	"web3/app/job/databus/internal/client"
	"web3/app/job/databus/internal/conf"
	"web3/app/job/datawatch/storage/mq"
)

type BlockEth struct {
	RabbitMQConsume *rabbitmq.Consumer
}

func NewBlockEth() *BlockEth {
	return &BlockEth{}
}

func (t *BlockEth) Name() string {
	return "Consumer BlockAptos"
}
func (t *BlockEth) Start(ctx context.Context) error {
	consumeConf := conf.Conf.GetRabbitmq().GetCWeb3NewBlockEth()
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
func (t *BlockEth) Stop(ctx context.Context) error {
	t.RabbitMQConsume.Close()
	return nil
}
func (t *BlockEth) Reload(ctx context.Context) error {
	return nil
}
func (t *BlockEth) _syncBlockData(d rabbitmq.Delivery) rabbitmq.Action {
	ctx := context.Background()
	var err error
	msgData := &mq.EthChainBlockHeader{}
	err = json.Unmarshal(d.Body, msgData)
	if err != nil {
		log.Context(ctx).Errorf("[Consume] BlockAptos SyncBlockInfo err:%+v", err)
		return rabbitmq.Ack
	}
	if msgData.Header.Number == nil {
		return rabbitmq.Ack
	}
	blockInfo, err := client.GlobalClient.Web3ClientPool.GetClient(msgData.ChainID).BlockByNumber(ctx, big.NewInt(msgData.Header.Number.Int64()))
	if err != nil {
		log.Context(ctx).Errorf("[Consume] BlockAptos BlockByNumber ChainUniqueID:%+v ChainID:%+v Number:%+v err:%+v", msgData.ChainUniqueID, msgData.ChainID, msgData.Header.Number.Int64(), err)
		return rabbitmq.Ack
	}
	if blockInfo == nil {
		return rabbitmq.Ack
	}
	syncReq := &data.SyncBlockInfoRequest{Item: &data.BlockItem{
		Number:            blockInfo.Number().Int64(),
		Hash:              blockInfo.Hash().String(),
		ParentHash:        blockInfo.ParentHash().String(),
		Timestamp:         int64(blockInfo.Time()),
		Difficulty:        blockInfo.Difficulty().String(),
		ExtraData:         common.BytesToHash(blockInfo.Extra()).String(),
		GasLimit:          strconv.FormatUint(blockInfo.GasLimit(), 10),
		GasUsed:           strconv.FormatUint(blockInfo.GasUsed(), 10),
		BaseFeePerGas:     blockInfo.BaseFee().String(),
		Miner:             common.HexToAddress("").Hex(),
		MixHash:           common.HexToHash("").Hex(),
		Nonce:             strconv.FormatUint(blockInfo.Nonce(), 10),
		ReceiptsRoot:      blockInfo.Header().ReceiptHash.String(),
		Sha3Uncles:        common.HexToHash("").Hex(),
		StateRoot:         common.HexToHash("").Hex(),
		Size:              int64(blockInfo.Size()),
		TransactionsRoot:  blockInfo.Root().String(),
		TransactionsCount: int64(len(blockInfo.Transactions())),
		UnclesCount:       int64(len(blockInfo.Uncles())),
		ChainId:           msgData.ChainID,
		ChainUniqueId:     int64(msgData.ChainUniqueID),
	}}

	_, err = client.GlobalClient.ClientAppServiceDataChainRPC.SyncBlockInfo(ctx, syncReq)
	if err != nil {
		log.Context(ctx).Errorf("[Consume] BlockAptos SyncBlockInfo err:%+v", err)
		return rabbitmq.NackRequeue
	}

	for idx, tx := range blockInfo.Transactions() {
		fromAddress := common.HexToAddress("").String()
		toAddress := common.HexToAddress("").String()
		if tx.To() != nil {
			toAddress = tx.To().String()
		}
		r, s, v := tx.RawSignatureValues()
		syncTxReq := &data.SyncBlockTransactionInfoRequest{
			Item: &data.TransactionItem{
				Type:                 int64(tx.Type()),
				Status:               1,
				BlockHash:            blockInfo.Hash().String(),
				BlockNumber:          blockInfo.Number().Int64(),
				BlockTimestamp:       int64(blockInfo.Time()),
				TransactionHash:      tx.Hash().String(),
				TransactionIndex:     int64(idx),
				FromAddress:          fromAddress,
				ToAddress:            toAddress,
				Value:                tx.Value().String(),
				Input:                crypto.Keccak256Hash(tx.Data()).Hex(),
				Nonce:                int64(tx.Nonce()),
				ContractAddress:      common.HexToAddress("").String(),
				Gas:                  int64(tx.Gas()),
				GasPrice:             tx.GasPrice().Int64(),
				GasUsed:              0,
				EffectiveGasPrice:    0,
				CumulativeGasUsed:    0,
				MaxFeePerGas:         0,
				MaxPriorityFeePerGas: 0,
				R:                    r.String(),
				S:                    s.String(),
				V:                    v.Int64(),
				LogsCount:            0,
				ChainId:              msgData.ChainID,
				ChainUniqueId:        int64(msgData.ChainUniqueID),
			},
		}

		syncTxResp, loopErr := client.GlobalClient.ClientAppServiceDataChainRPC.SyncBlockTransactionInfo(ctx, syncTxReq)
		if err != nil {
			log.Errorf("SyncTransactionInfo saveReq: %+v syncTxReq: %+v err: %+v err: %+v", syncTxReq, syncTxResp, loopErr)
			return rabbitmq.NackRequeue
		}
	}
	return rabbitmq.Ack
}
