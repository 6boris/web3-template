package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	internalClient "web3/app/job/datawatch/internal/client"
	"web3/app/job/datawatch/internal/dao"
	"web3/app/job/datawatch/storage/mq"
)

type WatchAptosNewBlockHeader struct {
	dataCh       chan *ConsumerData
	stopNotice   chan struct{}
	Hooks        []ConsumerHook
	hooks        []ConsumerHook
	watchHooks   map[string]WatchHook
	consumeHooks map[string]ConsumerHook
}

func NewWatchAptosNewBlockHeader() *WatchAptosNewBlockHeader {
	inst := &WatchAptosNewBlockHeader{
		stopNotice: make(chan struct{}, 2),
		dataCh:     make(chan *ConsumerData, 10),
	}
	inst.watchHooks = map[string]WatchHook{
		"_watchAptosMainnet": inst._watchAptosMainnet,
		"_watchAptosTestnet": inst._watchAptosTestnet,
	}
	inst.consumeHooks = map[string]ConsumerHook{
		"_consumeSendDataToRabbitMQ": inst._consumeSendDataToRabbitMQ,
	}

	return inst
}
func (t *WatchAptosNewBlockHeader) Name() string {
	return "WatchAptosNewBlockHeader Task"
}
func (t *WatchAptosNewBlockHeader) Start(ctx context.Context) error {
	go func() {
		_ = t.watch(ctx)
	}()
	go func() {
		t.consume()
	}()
	return nil
}
func (t *WatchAptosNewBlockHeader) Stop(ctx context.Context) error {
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	return nil
}
func (t *WatchAptosNewBlockHeader) Reload(ctx context.Context) error {
	return nil
}

func (t *WatchAptosNewBlockHeader) watch(ctx context.Context) error {
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
func (t *WatchAptosNewBlockHeader) consume() {
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
func (t *WatchAptosNewBlockHeader) _watchAptosMainnet(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(1)
	nodeInfo, err := internalClient.GlobalClient.ClientAptosGetNodeInfo(ctx, chainID)
	if err != nil {
		log.Fatal(err)
	}
	latestBlockHeight := nodeInfo.BlockHeight
	for {
		select {
		case <-t.stopNotice:
			log.Warnf("Stop watch WatchAptosNewBlockHeader _watchAptosMainnet: %+v", t.Name())
			return nil
		case <-time.Tick(time.Second):
			blockInfo, err := internalClient.GlobalClient.ClientAptosGetBlocksByHeight(ctx, chainID, latestBlockHeight)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			hByte, _ := json.Marshal(blockInfo)
			ch <- &ConsumerData{RawData: string(hByte)}
			latestBlockHeight++
		}
	}
}
func (t *WatchAptosNewBlockHeader) _watchAptosTestnet(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(2)
	nodeInfo, err := internalClient.GlobalClient.ClientAptosGetNodeInfo(ctx, chainID)
	if err != nil {
		log.Fatal(err)
	}
	latestBlockHeight := nodeInfo.BlockHeight
	for {
		select {
		case <-t.stopNotice:
			log.Warnf("Stop watch WatchAptosNewBlockHeader _watchAptosTestnet: %+v", t.Name())
			return nil
		case <-time.Tick(time.Second):
			blockInfo, err := internalClient.GlobalClient.ClientAptosGetBlocksByHeight(ctx, chainID, latestBlockHeight)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			hByte, _ := json.Marshal(blockInfo)
			ch <- &ConsumerData{RawData: string(hByte)}
			latestBlockHeight++
		}
	}
}
func (t *WatchAptosNewBlockHeader) _consumeSendDataToRabbitMQ(ctx context.Context, data ConsumerData) error {
	blockHeader := &mq.AptosChainBlockInfo{}
	err := json.Unmarshal([]byte(data.RawData), blockHeader)
	if err != nil {
		return err
	}
	err = dao.GlobalDao.RabbitMQPushWeb3EAptosNewBlockMsg(ctx, blockHeader)
	log.Infof("WatchAptosNewBlockHeader _consumeSendDataToRabbitMQ ChainUniqueID:%+v ChainID:%+v Number:%+v err:%+v", blockHeader.ChainUniqueID, blockHeader.ChainID, blockHeader.BlockHeight, err)
	return err
}
