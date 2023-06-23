package client

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"net/http"
	"web3/app/service/data/storage/model"
)

func (c *Client) DefiGetTokenPrice(ctx context.Context, sourceToken, targetToken string) (*model.DefiPrice, error) {
	data := &model.DefiPrice{
		SourceToken: sourceToken,
		TargetToken: targetToken,
	}
	resp, err := req.SetContext(ctx).Get(fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s", sourceToken, targetToken))
	if err != nil {
		return nil, err
	}
	if gjson.Get(resp.String(), "Response").String() != "" {
		return nil, errors.New(http.StatusTooManyRequests, "RATE_LIMIT", gjson.Get(resp.String(), "Message").String())
	}

	data.Price = gjson.Get(resp.String(), fmt.Sprintf("RAW.%s.%s.PRICE", sourceToken, targetToken)).Float()
	data.Timestamp = gjson.Get(resp.String(), fmt.Sprintf("RAW.%s.%s.CONVERSIONLASTUPDATE", sourceToken, targetToken)).Int()
	return data, nil
}
