package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"math/big"
	"strings"
	"time"
	"web3/api/base"
	internalClient "web3/app/job/datawatch/internal/client"
	"web3/app/job/datawatch/internal/dao"
	"web3/app/job/datawatch/storage/mq"
)

type WatchPolygonNewBlockHeader struct {
	dataCh       chan *ConsumerData
	stopNotice   chan struct{}
	Hooks        []ConsumerHook
	hooks        []ConsumerHook
	watchHooks   map[string]WatchHook
	consumeHooks map[string]ConsumerHook
}

func NewWatchPolygonNewBlockHeader() *WatchPolygonNewBlockHeader {
	inst := &WatchPolygonNewBlockHeader{
		stopNotice: make(chan struct{}, 2),
		dataCh:     make(chan *ConsumerData, 10),
	}
	inst.watchHooks = map[string]WatchHook{
		"_watchPolygonMainnet": inst._watchPolygonMainnet,
		"_watchPolygonMumbai":  inst._watchPolygonMumbai,
	}
	inst.consumeHooks = map[string]ConsumerHook{
		"_consumeSendDataToRabbitMQ": inst._consumeSendDataToRabbitMQ,
	}

	return inst
}
func (t *WatchPolygonNewBlockHeader) Name() string {
	return "WatchPolygonNewBlockHeader Task"
}
func (t *WatchPolygonNewBlockHeader) Start(ctx context.Context) error {
	go func() {
		_ = t.watch(ctx)
	}()
	go func() {
		t.consume()
	}()
	return nil
}
func (t *WatchPolygonNewBlockHeader) Stop(ctx context.Context) error {
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	return nil
}
func (t *WatchPolygonNewBlockHeader) Reload(ctx context.Context) error {
	return nil
}

func (t *WatchPolygonNewBlockHeader) watch(ctx context.Context) error {
	for tmpK, tmpF := range t.watchHooks {
		f := tmpF
		k := tmpK
		go func() {
			tmpErr := f(context.Background(), t.dataCh)
			fmt.Println("Start watch", t.Name(), k, time.Now(), tmpErr)
		}()
	}
	return nil
}
func (t *WatchPolygonNewBlockHeader) consume() {
	fmt.Println("Start consume", t.Name(), time.Now())
	for {
		select {
		case l := <-t.dataCh:
			for _, f := range t.consumeHooks {
				hookErr := f(context.Background(), *l)
				if hookErr != nil {
					fmt.Println(hookErr)
				}
			}
		}
	}
}
func (t *WatchPolygonNewBlockHeader) _watchPolygonMainnet(ctx context.Context, ch chan *ConsumerData) error {
	client, err := ethclient.Dial("wss://polygon-mainnet.s.chainbase.online/v1/2RJgEIAYv2NjIcwCfEwQ54aAbYR")
	if err != nil {
		return err
	}
	chainSub := make(chan *types.Header, 10)
	_, err = client.SubscribeNewHead(context.Background(), chainSub)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-t.stopNotice:
				fmt.Println("Stop watch", t.Name(), time.Now())
				return
			case l := <-chainSub:
				headerInfo := &mq.EthChainBlockHeader{
					StreamID:      fmt.Sprintf("%s:_watchPolygonMainnet:%d:%s", t.Name(), base.CHAIN_UNIQUE_ID_POLYGON_MAINNET, strings.ReplaceAll(uuid.NewString(), "-", "")),
					ChainID:       137,
					ChainUniqueID: base.CHAIN_UNIQUE_ID_POLYGON_MAINNET,
					Header:        l,
				}
				hByte, _ := json.Marshal(headerInfo)
				ch <- &ConsumerData{RawData: string(hByte)}
			}
		}
	}()
	return nil
}
func (t *WatchPolygonNewBlockHeader) _watchPolygonMumbai(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(80001)
	latestBlockNumber, err := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockNumber(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	for {
		select {
		case <-t.stopNotice:
			log.Warnf("Stop watch WatchEthereumNewBlockHeader _watchEthereumSepolia: %+v", t.Name())
			return nil
		case <-time.Tick(time.Second):
			blockInfo, loopErr := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockByNumber(ctx, big.NewInt(int64(latestBlockNumber)))
			if loopErr != nil {
				continue
			}
			headerInfo := &mq.EthChainBlockHeader{
				StreamID:      fmt.Sprintf("%s:_watchEthereumSepolia:%d:%s", t.Name(), base.CHAIN_UNIQUE_ID_ETHEREUM_SEPOLIA, strings.ReplaceAll(uuid.NewString(), "-", "")),
				ChainID:       chainID,
				ChainUniqueID: base.CHAIN_UNIQUE_ID_ETHEREUM_SEPOLIA,
				Header:        blockInfo.Header(),
			}
			hByte, _ := json.Marshal(headerInfo)
			ch <- &ConsumerData{RawData: string(hByte)}
			latestBlockNumber++
		}
	}
}
func (t *WatchPolygonNewBlockHeader) _consumeSendDataToRabbitMQ(ctx context.Context, data ConsumerData) error {
	blockHeader := &mq.EthChainBlockHeader{}
	err := json.Unmarshal([]byte(data.RawData), blockHeader)
	if err != nil {
		return err
	}
	err = dao.GlobalDao.RabbitMQPushWeb3EthNewBlockMsg(ctx, blockHeader)
	log.Infof("WatchFantomNewBlockHeader _consumeSendDataToRabbitMQ ChainUniqueID:%+v ChainID:%+v Number:%+v err:%+v", blockHeader.ChainUniqueID, blockHeader.ChainID, blockHeader.Header.Number, err)
	return err
}
