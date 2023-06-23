package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httpReq "github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"web3/api/base"
	"web3/app/job/datawatch/storage/mq"
)

var APTOS_RPC_URLS = map[int64]string{
	1: "https://fullnode.mainnet.aptoslabs.com",
	2: "https://fullnode.testnet.aptoslabs.com",
}
var APTOS_UNIQUE_ID = map[int64]base.CHAIN_UNIQUE_ID{
	1: base.CHAIN_UNIQUE_ID_APTOS_MAINNET,
	2: base.CHAIN_UNIQUE_ID_APTOS_TESTNET,
}

func (c *Client) ClientAptosGetNodeInfo(ctx context.Context, chainID int64) (*mq.AptosChainNodeInfo, error) {
	url := ""
	if v, ok := APTOS_RPC_URLS[chainID]; ok {
		url = v
	}
	if url == "" {
		return nil, errors.New("not support this chain")
	}
	url = fmt.Sprintf("%s/v1", url)
	data := &mq.AptosChainNodeInfo{}
	resp, err := httpReq.R().SetContext(ctx).Get(url)
	if err != nil {
		return data, nil
	}
	data.ChainID = gjson.Get(resp.String(), "chain_id").Int()
	data.Epoch = gjson.Get(resp.String(), "epoch").Int()
	data.LedgerVersion = gjson.Get(resp.String(), "ledger_version").Int()
	data.LedgerTimestamp = gjson.Get(resp.String(), "ledger_timestamp").Int()
	data.NodeRole = gjson.Get(resp.String(), "node_role").String()
	data.BlockHeight = gjson.Get(resp.String(), "block_height").Int()
	data.GitHash = gjson.Get(resp.String(), "git_hash").String()
	if v, ok := APTOS_UNIQUE_ID[chainID]; ok {
		data.ChainUniqueID = v
	}
	return data, nil
}

func (c *Client) ClientAptosGetBlocksByHeight(ctx context.Context, chainID int64, blockHeight int64) (*mq.AptosChainBlockInfo, error) {
	url := ""
	if v, ok := APTOS_RPC_URLS[chainID]; ok {
		url = v
	}
	if url == "" {
		return nil, errors.New("not support this chain")
	}
	url = fmt.Sprintf("%s/v1/blocks/by_height/%d?with_transactions=true", url, blockHeight)
	data := &mq.AptosChainBlockInfo{
		Transactions: []*mq.AptosChainBTransactionsInfo{},
	}
	resp, err := httpReq.R().SetContext(ctx).Get(url)
	if err != nil {
		return data, nil
	}
	data.ChainID = chainID
	data.BlockHeight = gjson.Get(resp.String(), "block_height").Int()
	if v, ok := APTOS_UNIQUE_ID[chainID]; ok {
		data.ChainUniqueID = v
	}
	data.BlockHash = gjson.Get(resp.String(), "block_hash").String()
	data.BlockTimestamp = gjson.Get(resp.String(), "block_timestamp").Int()
	data.FirstVersion = gjson.Get(resp.String(), "first_version").Int()
	data.LastVersion = gjson.Get(resp.String(), "last_version").Int()
	tx := make([]*mq.AptosChainBTransactionsInfo, 0)
	err = json.Unmarshal([]byte(gjson.Get(resp.String(), "transactions").String()), &tx)
	if err != nil {
		return nil, err
	}
	data.Transactions = tx
	return data, nil
}
