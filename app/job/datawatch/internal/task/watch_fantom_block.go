package task

import (
	"context"
	"encoding/json"
	"fmt"
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

type WatchFantomNewBlockHeader struct {
	dataCh       chan *ConsumerData
	stopNotice   chan struct{}
	Hooks        []ConsumerHook
	hooks        []ConsumerHook
	watchHooks   map[string]WatchHook
	consumeHooks map[string]ConsumerHook
}

func NewWatchFantomNewBlockHeader() *WatchFantomNewBlockHeader {
	inst := &WatchFantomNewBlockHeader{
		stopNotice: make(chan struct{}, 1),
		dataCh:     make(chan *ConsumerData, 10),
	}
	inst.watchHooks = map[string]WatchHook{
		"_watchFantomMainnet": inst._watchFantomMainnet,
		"_watchFantomTestnet": inst._watchFantomTestnet,
	}
	inst.consumeHooks = map[string]ConsumerHook{
		"_consumeSendDataToRabbitMQ": inst._consumeSendDataToRabbitMQ,
	}

	return inst
}
func (t *WatchFantomNewBlockHeader) Name() string {
	return "WatchFantomNewBlockHeader Task"
}
func (t *WatchFantomNewBlockHeader) Start(ctx context.Context) error {
	go func() {
		_ = t.watch(ctx)
	}()
	go func() {
		t.consume()
	}()
	return nil
}
func (t *WatchFantomNewBlockHeader) Stop(ctx context.Context) error {
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	return nil
}
func (t *WatchFantomNewBlockHeader) Reload(ctx context.Context) error {
	return nil
}

func (t *WatchFantomNewBlockHeader) watch(ctx context.Context) error {
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
func (t *WatchFantomNewBlockHeader) consume() {
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
func (t *WatchFantomNewBlockHeader) _watchFantomMainnet(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(250)
	latestBlockNumber, err := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockNumber(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	for {
		select {
		case <-t.stopNotice:
			log.Warnf("Stop watch WatchEthereumNewBlockHeader _watchFantomMainnet: %+v", t.Name())
			return nil
		case <-time.Tick(time.Second):
			blockInfo, loopErr := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockByNumber(ctx, big.NewInt(int64(latestBlockNumber)))
			if loopErr != nil {
				continue
			}
			headerInfo := &mq.EthChainBlockHeader{
				StreamID:      fmt.Sprintf("%s:_watchFantomMainnet:%d:%s", t.Name(), base.CHAIN_UNIQUE_ID_FANTOM_MAINNET, strings.ReplaceAll(uuid.NewString(), "-", "")),
				ChainID:       chainID,
				ChainUniqueID: base.CHAIN_UNIQUE_ID_FANTOM_MAINNET,
				Header:        blockInfo.Header(),
			}
			hByte, _ := json.Marshal(headerInfo)
			ch <- &ConsumerData{RawData: string(hByte)}
			latestBlockNumber++
		}
	}
}
func (t *WatchFantomNewBlockHeader) _watchFantomTestnet(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(4002)
	latestBlockNumber, err := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockNumber(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	for {
		select {
		case <-t.stopNotice:
			log.Warnf("Stop watch WatchEthereumNewBlockHeader _watchFantomTestnet: %+v", t.Name())
			return nil
		case <-time.Tick(time.Second):
			blockInfo, loopErr := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockByNumber(ctx, big.NewInt(int64(latestBlockNumber)))
			if loopErr != nil {
				continue
			}
			headerInfo := &mq.EthChainBlockHeader{
				StreamID:      fmt.Sprintf("%s:_watchFantomTestnet:%d:%s", t.Name(), base.CHAIN_UNIQUE_ID_FANTOM_TESTNET, strings.ReplaceAll(uuid.NewString(), "-", "")),
				ChainID:       chainID,
				ChainUniqueID: base.CHAIN_UNIQUE_ID_FANTOM_TESTNET,
				Header:        blockInfo.Header(),
			}
			hByte, _ := json.Marshal(headerInfo)
			ch <- &ConsumerData{RawData: string(hByte)}
			latestBlockNumber++
		}
	}
}
func (t *WatchFantomNewBlockHeader) _consumeSendDataToRabbitMQ(ctx context.Context, data ConsumerData) error {
	blockHeader := &mq.EthChainBlockHeader{}
	err := json.Unmarshal([]byte(data.RawData), blockHeader)
	if err != nil {
		return err
	}
	err = dao.GlobalDao.RabbitMQPushWeb3EthNewBlockMsg(ctx, blockHeader)
	log.Infof("WatchFantomNewBlockHeader _consumeSendDataToRabbitMQ ChainUniqueID:%+v ChainID:%+v Number:%+v err:%+v", blockHeader.ChainUniqueID, blockHeader.ChainID, blockHeader.Header.Number, err)
	return err
}
