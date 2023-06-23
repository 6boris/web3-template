package worker

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"sync"
	"time"
	"web3/app/job/analysis/internal/client"
	"web3/app/job/analysis/internal/dao"
	"web3/app/job/analysis/storage/mysql/model"
)

func (s *Server) _metricsWeb3ChainGasPriceCronSyncData() {
	ctx := context.Background()
	chainList, err := dao.GlobalDao.MySQLSearchChainInfoList(ctx, []int64{
		1000100,
		1000300,
		1000400,
		1000500,
		1000601,
		1000101,
		1000102,
		1000301,
		1000501,
		1001400,
		1001401,
		1000700,
		1000701,
	})
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for _, tmp := range chainList {
		wg.Add(1)
		chain := tmp
		if chain.ChainID == 0 {
			continue
		}
		go func() {
			gasPrice, loopErr := client.GlobalClient.Web3ClientPool.GetClient(chain.ChainID).SuggestGasPrice(ctx)
			if loopErr != nil {
				log.Context(ctx).Errorf("_metricsWeb3ChainGasPriceCronSyncData SuggestGasPrice UniqueID:%+v ChainID:%+v err:%+v", chain.UniqueID, chain.ChainID, loopErr)
				return
			}
			metricsData := &model.ChainMetric{
				ChainUniqueID: chain.UniqueID,
			}
			blockNumber, loopErr := client.GlobalClient.Web3ClientPool.GetClient(chain.ChainID).BlockNumber(ctx)
			if loopErr != nil {
				log.Context(ctx).Errorf("_metricsWeb3ChainGasPriceCronSyncData BlockNumber UniqueID:%+v ChainID:%+v err:%+v", chain.UniqueID, chain.ChainID, loopErr)
				return
			}
			s._metricsWeb3ChainGasPriceGaugeLocalData.Store(chain.UniqueID, gasPrice.Int64())
			metricsData.LatestBlockNumber = int64(blockNumber)
			metricsData.GasPrice = gasPrice.Int64()
			err = dao.GlobalDao.MySQLSaveChainMetric(ctx, metricsData)

		}()
	}
	wg.Wait()
}
func (s *Server) _metricsWeb3AptosLatestBlockNumberCronSyncData() {
	ctx := context.Background()
	chainList, err := dao.GlobalDao.MySQLSearchChainInfoList(ctx, []int64{
		1003500,
		1003501,
	})
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for _, tmp := range chainList {
		wg.Add(1)
		chain := tmp
		if chain.ChainID == 0 {
			continue
		}
		go func() {
			metricsData := &model.ChainMetric{
				ChainUniqueID: chain.UniqueID,
			}
			blockNumber, loopErr := client.GlobalClient.ClientAptosGetNodeInfo(ctx, chain.ChainID)
			if loopErr != nil {
				return
			}
			metricsData.LatestBlockNumber = blockNumber.BlockHeight
			err = dao.GlobalDao.MySQLSaveChainMetric(ctx, metricsData)

		}()
	}
	wg.Wait()
}
func (s *Server) _metricsWeb3ChainGasPriceCallback(ctx context.Context, o metric.Int64Observer) error {
	s._metricsWeb3ChainGasPriceGaugeLocalData.Range(func(key, value any) bool {
		chainName := ""
		chainEnv := ""
		chainID := int64(0)
		chainUniqueID := key.(int64)
		if v, ok := s._localCacheChainInfo.Load(chainUniqueID); ok {
			chain := v.(*model.ChainInfo)
			chainName = chain.ChainName
			chainEnv = chain.ChainEnv
			chainID = chain.ChainID
		} else {
			return true
		}
		o.Observe(value.(int64),
			metric.WithAttributes(
				attribute.Key("chain_unique_id").Int64(chainUniqueID),
				attribute.Key("chain_id").Int64(chainID),
				attribute.Key("chain_name").String(chainName),
				attribute.Key("chain_env").String(chainEnv),
			),
		)
		return true
	})
	return nil
}
func (s *Server) _metricsWeb3ChainTransactionTpsCronSyncData() {
	ctx := context.Background()
	chainList, err := dao.GlobalDao.MySQLSearchChainInfoList(ctx, []int64{
		1000100,
		1000300,
		1000400,
		1000500,
		1000601,
		1000101,
		1000102,
		1000301,
		1000501,
		1001400,
		1001401,
		1000700,
		1000701,
		1003500,
		1003501,
	})
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for _, tmp := range chainList {
		wg.Add(1)
		chain := tmp
		if chain.ChainID == 0 {
			continue
		}
		go func() {
			metricsData := &model.ChainMetric{
				ChainUniqueID: chain.UniqueID,
			}
			txTps, loopErr := dao.GlobalDao.MySQLGetRecentTxMetrics(ctx, chain.UniqueID, time.Now().Add(-time.Minute*2), time.Now().Add(-time.Minute*1), 60)
			if loopErr != nil {
				return
			}
			s._metricsWeb3ChainTransactionTpsGaugeLocalData.Store(chain.UniqueID, txTps)
			metricsData.TxTps = txTps
			err = dao.GlobalDao.MySQLSaveChainMetric(ctx, metricsData)

		}()
	}
	wg.Wait()
}
func (s *Server) _metricsWeb3ChainTransactionTpdCronSyncData() {
	ctx := context.Background()
	chainList, err := dao.GlobalDao.MySQLSearchChainInfoList(ctx, []int64{
		1000100,
		1000300,
		1000400,
		1000500,
		1000601,
		1000101,
		1000102,
		1000301,
		1000501,
		1001400,
		1001401,
		1000700,
		1000701,
		1003500,
		1003501,
	})
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for _, tmp := range chainList {
		wg.Add(1)
		chain := tmp
		if chain.ChainID == 0 {
			continue
		}
		go func() {
			metricsData := &model.ChainMetric{
				ChainUniqueID: chain.UniqueID,
			}
			txTpd, loopErr := dao.GlobalDao.MySQLGetRecentTxMetrics(ctx, chain.UniqueID, time.Now().Add(-time.Hour*24), time.Now(), 1)
			if loopErr != nil {
				return
			}
			metricsData.TxTpd = txTpd
			err = dao.GlobalDao.MySQLSaveChainMetric(ctx, metricsData)

		}()
	}
	wg.Wait()
}
func (s *Server) _metricsWeb3ChainTransactionTpsCallback(ctx context.Context, o metric.Float64Observer) error {
	s._metricsWeb3ChainTransactionTpsGaugeLocalData.Range(func(key, value any) bool {
		chainName := ""
		chainEnv := ""
		chainID := int64(0)
		chainUniqueID := key.(int64)
		if v, ok := s._localCacheChainInfo.Load(chainUniqueID); ok {
			chain := v.(*model.ChainInfo)
			chainName = chain.ChainName
			chainEnv = chain.ChainEnv
			chainID = chain.ChainID
		} else {
			return true
		}
		o.Observe(value.(float64),
			metric.WithAttributes(
				attribute.Key("chain_unique_id").Int64(chainUniqueID),
				attribute.Key("chain_id").Int64(chainID),
				attribute.Key("chain_name").String(chainName),
				attribute.Key("chain_env").String(chainEnv),
			),
		)
		return true
	})
	return nil
}

func (s *Server) _metricsWeb3ChainTransactionPendingCronSyncData() {
	ctx := context.Background()
	chainList, err := dao.GlobalDao.MySQLSearchChainInfoList(ctx, []int64{
		1000100,
		1000300,
		1000400,
		1000500,
		1000601,
		1000101,
		1000102,
		1000301,
		1000501,
		1001400,
		1001401,
		1000700,
		1000701,
	})
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for _, tmp := range chainList {
		wg.Add(1)
		chain := tmp
		if chain.ChainID == 0 {
			continue
		}
		go func() {
			metricsData := &model.ChainMetric{
				ChainUniqueID: chain.UniqueID,
			}
			txPending, loopErr := client.GlobalClient.Web3ClientPool.GetClient(chain.ChainID).PendingTransactionCount(ctx)
			if loopErr != nil {
				return
			}
			metricsData.TxPending = int64(txPending)
			err = dao.GlobalDao.MySQLSaveChainMetric(ctx, metricsData)
			s._metricsWeb3ChainTransactionPendingGaugeLocalData.Store(chain.UniqueID, int64(txPending))
		}()
	}
	wg.Wait()
}
func (s *Server) _metricsWeb3ChainTransactionPendingCallback(ctx context.Context, o metric.Int64Observer) error {
	s._metricsWeb3ChainTransactionPendingGaugeLocalData.Range(func(key, value any) bool {
		chainName := ""
		chainEnv := ""
		chainID := int64(0)
		chainUniqueID := key.(int64)
		if v, ok := s._localCacheChainInfo.Load(chainUniqueID); ok {
			chain := v.(*model.ChainInfo)
			chainName = chain.ChainName
			chainEnv = chain.ChainEnv
			chainID = chain.ChainID
		} else {
			return true
		}
		o.Observe(value.(int64),
			metric.WithAttributes(
				attribute.Key("chain_unique_id").Int64(chainUniqueID),
				attribute.Key("chain_id").Int64(chainID),
				attribute.Key("chain_name").String(chainName),
				attribute.Key("chain_env").String(chainEnv),
			),
		)
		return true
	})
	return nil
}

func (s *Server) _metricsWeb3ChainTokenPriceUsdCronSyncData() {
	ctx := context.Background()
	chainList, err := dao.GlobalDao.MySQLSearchChainInfoList(ctx, []int64{
		1000100,
		1000500,
		1000300,
		1000700,
		1001400,
		1000101,
		1000102,
		1000301,
		1000701,
		1001401,
		1000501,
	})
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for _, tmp := range chainList {
		wg.Add(1)
		chain := tmp
		if chain.ChainID == 0 {
			continue
		}

		if chain.TokenSymbol == "" {
			continue
		}
		go func() {
			metricsData := &model.ChainMetric{
				ChainUniqueID: chain.UniqueID,
			}
			tokenPrice, loopErr := client.GlobalClient.DefiGetPrice(ctx, chain.TokenSymbol, "USD")
			if loopErr != nil {
				return
			}
			metricsData.TokenPriceUsd = tokenPrice.GetPrice()
			err = dao.GlobalDao.MySQLSaveChainMetric(ctx, metricsData)
			s._metricsWeb3ChainTokenPriceUsdGaugeLocalData.Store(chain.UniqueID, tokenPrice.GetPrice())
		}()
	}
	wg.Wait()
}
func (s *Server) _metricsWeb3ChainTokenPriceUsdCallback(ctx context.Context, o metric.Float64Observer) error {
	s._metricsWeb3ChainTokenPriceUsdGaugeLocalData.Range(func(key, value any) bool {
		chainName := ""
		chainEnv := ""
		chainID := int64(0)
		chainUniqueID := key.(int64)
		if v, ok := s._localCacheChainInfo.Load(chainUniqueID); ok {
			chain := v.(*model.ChainInfo)
			chainName = chain.ChainName
			chainEnv = chain.ChainEnv
			chainID = chain.ChainID
		} else {
			return true
		}
		o.Observe(value.(float64),
			metric.WithAttributes(
				attribute.Key("chain_unique_id").Int64(chainUniqueID),
				attribute.Key("chain_id").Int64(chainID),
				attribute.Key("chain_name").String(chainName),
				attribute.Key("chain_env").String(chainEnv),
			),
		)
		return true
	})
	return nil
}
