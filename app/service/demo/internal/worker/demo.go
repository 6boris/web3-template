package worker

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"go.opentelemetry.io/otel/trace"
	"math/rand"
	"strings"
	"time"
	"web3/api/service/demo"
	"web3/app/service/demo/internal/client"
	"web3/app/service/demo/internal/conf"
	"web3/app/service/demo/internal/dao"
	"web3/pkg/otel"
)

func (s *Server) _demoDefiPrice() {
	tokenList := []string{
		"BTC", "ETH", "USDT", "BNB", "USDC", "XRP", "STETH", "ADA", "DOGE", "SOL", "TRX", "MATIC", "LTC", "DOT", "WBTC", "SHIB", "AVAX", "DAI", "BUSD", "UNI", "LEO", "TUSD", "LINK", "OKB", "BCH", "ATOM", "XMR", "ETC", "XLM", "TON", "ICP", "LDO", "FIL", "HBAR", "APT", "QNT", "CRO", "ARB", "NEAR", "VET", "STX", "GRT", "FRAX", "USDP", "RETH", "OP", "ALGO", "APE", "RNDR", "EGLD", "FTM", "AAVE", "RPL", "SAND", "EOS", "IMX", "XTZ", "THETA", "USDD", "WBT", "MANA", "BGB", "XRD", "MKR", "KCS", "GALA", "NEO", "AXS", "PEPE", "SNX", "BIT", "FLOW", "CRV", "GT", "INJ", "GUSD", "KAVA", "BSV", "LUNC", "KLAY", "CFX", "CETH", "GMX", "PAXG", "MIOTA", "TKX", "XAUT", "KAS", "CSPR", "BTT", "XDC", "SUI", "MINA", "XEC", "FXS", "FRXETH", "HT", "CHZ", "TWT", "DASH",
	}
	rand.NewSource(time.Now().UnixMicro())
	for i := 0; i < 2; i++ {
		sourceToken := tokenList[rand.Int63n(int64(len(tokenList)))]
		targetToken := "USD"
		url := fmt.Sprintf("http://%s/api/gateway/interface/web3/defi/price?source_token=%s&target_token=%s",
			conf.Conf.GetClient().GetAppInterfaceWeb3Http(), sourceToken, targetToken)
		_, err := req.Get(url)
		if err != nil {
			return
		}
	}
}

func (s *Server) _demoClient() {
	traceId, spanID := otel.NewTraceIDs(context.Background())
	ctx := trace.ContextWithSpanContext(context.Background(), trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
		SpanID:  spanID,
	}))
	dataGRPC, err := client.GlobalClient.ClientAppServiceDemoUserGRPC.GetUser(ctx, &demo.GetUserRequest{Id: 1})
	if err != nil {
		log.Context(ctx).Errorf("_demoClient GetUser data:%+v err:%+v", dataGRPC, err)
		return
	}
	dataHTTP, err := client.GlobalClient.ClientAppServiceDemoUserHTTP.GetUser(ctx, &demo.GetUserRequest{Id: 1})
	if err != nil {
		log.Context(ctx).Errorf("_demoClient GetUser data:%+v err:%+v", dataHTTP, err)
		return
	}
}
func (s *Server) _demoDao() {
	traceId, _ := trace.TraceIDFromHex(strings.ReplaceAll(uuid.NewString(), "-", ""))

	ctx := trace.ContextWithSpanContext(context.Background(), trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	}))
	err := dao.GlobalDao.RedisDelUserInfo(ctx, 1)
	if err != nil {
		log.Context(ctx).Errorf("_demoDao RedisDelUserInfo err:%+v", err)
		return
	}
}
