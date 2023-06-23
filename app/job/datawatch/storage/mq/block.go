package mq

import (
	"github.com/ethereum/go-ethereum/core/types"
	"web3/api/base"
)

type EthChainBlockHeader struct {
	StreamID      string               `json:"stream_id"`
	ChainID       int64                `json:"chain_id"`
	ChainUniqueID base.CHAIN_UNIQUE_ID `json:"chain_unique_id"`
	Header        *types.Header        `json:"header"`
}

type AptosChainNodeInfo struct {
	ChainID         int64                `json:"chain_id"`
	ChainUniqueID   base.CHAIN_UNIQUE_ID `json:"chain_unique_id"`
	Epoch           int64                `json:"epoch"`
	LedgerVersion   int64                `json:"ledger_version"`
	LedgerTimestamp int64                `json:"ledger_timestamp"`
	NodeRole        string               `json:"node_role"`
	BlockHeight     int64                `json:"block_height"`
	GitHash         string               `json:"git_hash"`
}

type AptosChainBlockInfo struct {
	StreamID       string                         `json:"stream_id"`
	ChainID        int64                          `json:"chain_id"`
	ChainUniqueID  base.CHAIN_UNIQUE_ID           `json:"chain_unique_id"`
	BlockHeight    int64                          `json:"block_height"`
	BlockHash      string                         `json:"block_hash"`
	BlockTimestamp int64                          `json:"block_timestamp"`
	FirstVersion   int64                          `json:"first_version"`
	LastVersion    int64                          `json:"last_version"`
	Transactions   []*AptosChainBTransactionsInfo `json:"transactions"`
}

type AptosChainBTransactionsInfo struct {
	Version             string `json:"version"`
	Hash                string `json:"hash"`
	StateChangeHash     string `json:"state_change_hash"`
	EventRootHash       string `json:"event_root_hash"`
	StateCheckpointHash any    `json:"state_checkpoint_hash"`
	GasUsed             string `json:"gas_used"`
	Success             bool   `json:"success"`
	VMStatus            string `json:"vm_status"`
	AccumulatorRootHash string `json:"accumulator_root_hash"`
	Changes             []struct {
		Address      string `json:"address"`
		StateKeyHash string `json:"state_key_hash"`
		Data         struct {
			Type string `json:"type"`
			Data struct {
				EpochInterval  string `json:"epoch_interval"`
				Height         string `json:"height"`
				NewBlockEvents struct {
					Counter string `json:"counter"`
					GUID    struct {
						ID struct {
							Addr        string `json:"addr"`
							CreationNum string `json:"creation_num"`
						} `json:"id"`
					} `json:"guid"`
				} `json:"new_block_events"`
				UpdateEpochIntervalEvents struct {
					Counter string `json:"counter"`
					GUID    struct {
						ID struct {
							Addr        string `json:"addr"`
							CreationNum string `json:"creation_num"`
						} `json:"id"`
					} `json:"guid"`
				} `json:"update_epoch_interval_events"`
			} `json:"data"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"changes"`
	ID     string `json:"id,omitempty"`
	Epoch  string `json:"epoch,omitempty"`
	Round  string `json:"round,omitempty"`
	Events []struct {
		GUID struct {
			CreationNumber string `json:"creation_number"`
			AccountAddress string `json:"account_address"`
		} `json:"guid"`
		SequenceNumber string `json:"sequence_number"`
		Type           string `json:"type"`
		Data           struct {
			Epoch                    string `json:"epoch"`
			FailedProposerIndices    []any  `json:"failed_proposer_indices"`
			Hash                     string `json:"hash"`
			Height                   string `json:"height"`
			PreviousBlockVotesBitvec string `json:"previous_block_votes_bitvec"`
			Proposer                 string `json:"proposer"`
			Round                    string `json:"round"`
			TimeMicroseconds         string `json:"time_microseconds"`
		} `json:"data"`
	} `json:"events,omitempty"`
	PreviousBlockVotesBitvec []int  `json:"previous_block_votes_bitvec,omitempty"`
	Proposer                 string `json:"proposer,omitempty"`
	FailedProposerIndices    []any  `json:"failed_proposer_indices,omitempty"`
	Timestamp                string `json:"timestamp"`
	Type                     string `json:"type"`
}
