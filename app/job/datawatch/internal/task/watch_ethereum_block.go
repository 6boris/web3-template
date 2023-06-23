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

type WatchEthereumNewBlockHeader struct {
	dataCh       chan *ConsumerData
	stopNotice   chan struct{}
	Hooks        []ConsumerHook
	hooks        []ConsumerHook
	watchHooks   map[string]WatchHook
	consumeHooks map[string]ConsumerHook
}

func NewWatchEthereumNewBlockHeader() *WatchEthereumNewBlockHeader {
	inst := &WatchEthereumNewBlockHeader{
		stopNotice: make(chan struct{}, 1),
		dataCh:     make(chan *ConsumerData, 10),
	}
	inst.watchHooks = map[string]WatchHook{
		"_watchEthereumMainnet": inst._watchEthereumMainnet,
		"_watchEthereumGoerli":  inst._watchEthereumGoerli,
		"_watchEthereumSepolia": inst._watchEthereumSepolia,
	}
	inst.consumeHooks = map[string]ConsumerHook{
		"_consumeSendDataToRabbitMQ": inst._consumeSendDataToRabbitMQ,
	}

	return inst
}
func (t *WatchEthereumNewBlockHeader) Name() string {
	return "WatchEthereumNewBlockHeader Task"
}
func (t *WatchEthereumNewBlockHeader) Start(ctx context.Context) error {
	go func() {
		_ = t.watch(ctx)
	}()
	go func() {
		t.consume()
	}()
	return nil
}
func (t *WatchEthereumNewBlockHeader) Stop(ctx context.Context) error {
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	t.stopNotice <- struct{}{}
	return nil
}
func (t *WatchEthereumNewBlockHeader) Reload(ctx context.Context) error {
	return nil
}

func (t *WatchEthereumNewBlockHeader) watch(ctx context.Context) error {
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
func (t *WatchEthereumNewBlockHeader) consume() {
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
func (t *WatchEthereumNewBlockHeader) _watchEthereumMainnet(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(1)
	latestBlockNumber, err := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockNumber(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	for {
		select {
		case <-t.stopNotice:
			log.Warnf("Stop watch WatchEthereumNewBlockHeader _watchEthereumMainnet: %+v", t.Name())
			return nil
		case <-time.Tick(time.Second):
			blockInfo, loopErr := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockByNumber(ctx, big.NewInt(int64(latestBlockNumber)))
			if loopErr != nil {
				continue
			}
			headerInfo := &mq.EthChainBlockHeader{
				StreamID:      fmt.Sprintf("%s:_watchEthereumMainnet:%d:%s", t.Name(), base.CHAIN_UNIQUE_ID_ETHEREUM_MAINNET, strings.ReplaceAll(uuid.NewString(), "-", "")),
				ChainID:       chainID,
				ChainUniqueID: base.CHAIN_UNIQUE_ID_ETHEREUM_MAINNET,
				Header:        blockInfo.Header(),
			}
			hByte, _ := json.Marshal(headerInfo)
			ch <- &ConsumerData{RawData: string(hByte)}
			latestBlockNumber++
		}
	}
}
func (t *WatchEthereumNewBlockHeader) _watchEthereumGoerli(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(5)
	latestBlockNumber, err := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockNumber(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	for {
		select {
		case <-t.stopNotice:
			log.Warnf("Stop watch WatchEthereumNewBlockHeader _watchEthereumGoerli: %+v", t.Name())
			return nil
		case <-time.Tick(time.Second):
			blockInfo, loopErr := internalClient.GlobalClient.Web3ClientPool.GetClient(chainID).BlockByNumber(ctx, big.NewInt(int64(latestBlockNumber)))
			if loopErr != nil {
				continue
			}
			headerInfo := &mq.EthChainBlockHeader{
				StreamID:      fmt.Sprintf("%s:_watchEthereumGoerli:%d:%s", t.Name(), base.CHAIN_UNIQUE_ID_ETHEREUM_GOERLI, strings.ReplaceAll(uuid.NewString(), "-", "")),
				ChainID:       chainID,
				ChainUniqueID: base.CHAIN_UNIQUE_ID_ETHEREUM_GOERLI,
				Header:        blockInfo.Header(),
			}
			hByte, _ := json.Marshal(headerInfo)
			ch <- &ConsumerData{RawData: string(hByte)}
			latestBlockNumber++
		}
	}
}
func (t *WatchEthereumNewBlockHeader) _watchEthereumSepolia(ctx context.Context, ch chan *ConsumerData) error {
	chainID := int64(11155111)
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
func (t *WatchEthereumNewBlockHeader) _consumeSendDataToRabbitMQ(ctx context.Context, data ConsumerData) error {
	blockHeader := &mq.EthChainBlockHeader{}
	err := json.Unmarshal([]byte(data.RawData), blockHeader)
	if err != nil {
		return err
	}
	err = dao.GlobalDao.RabbitMQPushWeb3EthNewBlockMsg(ctx, blockHeader)
	log.Infof("WatchEthereumNewBlockHeader _consumeSendDataToRabbitMQ ChainUniqueID:%+v ChainID:%+v Number:%d err:%+v", blockHeader.ChainUniqueID, blockHeader.ChainID, blockHeader.Header.Number, err)
	return err
}
