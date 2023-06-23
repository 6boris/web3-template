package worker

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	otelProm "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	metricSdk "go.opentelemetry.io/otel/sdk/metric"
	"sync"
)

type ServerOption func(o *Server)

type Server struct {
	CronJob                                           *cron.Cron
	_localCacheChainInfo                              *sync.Map
	_metricsWeb3ChainGasPriceGaugeLocalData           *sync.Map
	_metricsWeb3ChainGasPriceGauge                    metric.Int64ObservableGauge
	_metricsWeb3ChainTransactionTpsGaugeLocalData     *sync.Map
	_metricsWeb3ChainTransactionTpsGauge              metric.Float64ObservableGauge
	_metricsWeb3ChainTransactionPendingGaugeLocalData *sync.Map
	_metricsWeb3ChainTransactionPendingGauge          metric.Int64ObservableGauge
	_metricsWeb3ChainTokenPriceUsdGaugeLocalData      *sync.Map
	_metricsWeb3ChainTokenPriceUsdGauge               metric.Float64ObservableGauge
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{}
	srv.CronJob = cron.New()
	exporter, err := otelProm.New(otelProm.WithoutTargetInfo())
	if err != nil {
		panic(err)
	}
	provider := metricSdk.NewMeterProvider(metricSdk.WithReader(exporter))
	meter := provider.Meter("Web3")

	srv._localCacheChainInfo = &sync.Map{}
	_, _ = srv.CronJob.AddFunc("@every 2s", srv._cronSyncChainInfo)
	_, _ = srv.CronJob.AddFunc("@every 2s", srv._cronSyncChainInfo)

	_, _ = srv.CronJob.AddFunc("@every 2s", srv._metricsWeb3AptosLatestBlockNumberCronSyncData)

	_, _ = srv.CronJob.AddFunc("@every 2s", srv._metricsWeb3ChainGasPriceCronSyncData)
	srv._metricsWeb3ChainGasPriceGaugeLocalData = &sync.Map{}
	srv._metricsWeb3ChainGasPriceGauge, err = meter.Int64ObservableGauge("web3_chain_gas_price", metric.WithDescription("Web3 Gas Price"), metric.WithInt64Callback(srv._metricsWeb3ChainGasPriceCallback))
	if err != nil {
		panic(err)
	}

	_, _ = srv.CronJob.AddFunc("@every 2s", srv._metricsWeb3ChainTransactionTpsCronSyncData)
	_, _ = srv.CronJob.AddFunc("@every 1m", srv._metricsWeb3ChainTransactionTpdCronSyncData)
	srv._metricsWeb3ChainTransactionTpsGaugeLocalData = &sync.Map{}
	srv._metricsWeb3ChainTransactionTpsGauge, err = meter.Float64ObservableGauge("web3_chain_transaction_tps", metric.WithDescription("Web3 Transaction Tps"), metric.WithFloat64Callback(srv._metricsWeb3ChainTransactionTpsCallback))
	if err != nil {
		panic(err)
	}

	_, _ = srv.CronJob.AddFunc("@every 2s", srv._metricsWeb3ChainTransactionPendingCronSyncData)
	srv._metricsWeb3ChainTransactionPendingGaugeLocalData = &sync.Map{}
	srv._metricsWeb3ChainTransactionPendingGauge, err = meter.Int64ObservableGauge("web3_chain_transaction_pending", metric.WithDescription("Web3 Transaction Tps"), metric.WithInt64Callback(srv._metricsWeb3ChainTransactionPendingCallback))
	if err != nil {
		panic(err)
	}

	_, _ = srv.CronJob.AddFunc("@every 2s", srv._metricsWeb3ChainTokenPriceUsdCronSyncData)
	srv._metricsWeb3ChainTokenPriceUsdGaugeLocalData = &sync.Map{}
	srv._metricsWeb3ChainTokenPriceUsdGauge, err = meter.Float64ObservableGauge("web3_chain_token_price_usd", metric.WithDescription("Web3 Transaction Price"), metric.WithFloat64Callback(srv._metricsWeb3ChainTokenPriceUsdCallback))
	if err != nil {
		panic(err)
	}

	return srv
}

func (s *Server) Start(ctx context.Context) error {
	log.Context(ctx).Info("[Worker] server starting")
	s.CronJob.Start()
	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	log.Context(ctx).Info("[Worker] server stopping")
	s.CronJob.Stop()
	return nil
}
