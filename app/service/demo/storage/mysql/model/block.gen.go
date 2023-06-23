// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBlock = "block"

// Block mapped from table <block>
type Block struct {
	ChainUniqueID     int64      `gorm:"column:chain_unique_id;type:bigint;primaryKey" json:"chain_unique_id"`
	Number            int64      `gorm:"column:number;type:bigint;primaryKey;comment:block number" json:"number"`
	Hash              string     `gorm:"column:hash;type:varchar(200);index:block_idx_hash,priority:1" json:"hash"`
	ParentHash        string     `gorm:"column:parent_hash;type:varchar(200)" json:"parent_hash"`
	Timestamp         *time.Time `gorm:"column:timestamp;type:datetime;not null;index:block_idx_timestamp,priority:1;default:CURRENT_TIMESTAMP;comment:Block Time" json:"timestamp"`
	Difficulty        string     `gorm:"column:difficulty;type:varchar(200)" json:"difficulty"`
	ExtraData         string     `gorm:"column:extra_data;type:text" json:"extra_data"`
	GasLimit          string     `gorm:"column:gas_limit;type:varchar(200)" json:"gas_limit"`
	GasUsed           string     `gorm:"column:gas_used;type:varchar(200)" json:"gas_used"`
	BaseFeePerGas     string     `gorm:"column:base_fee_per_gas;type:varchar(200)" json:"base_fee_per_gas"`
	Miner             string     `gorm:"column:miner;type:varchar(200)" json:"miner"`
	MixHash           string     `gorm:"column:mix_hash;type:varchar(200)" json:"mix_hash"`
	Nonce             string     `gorm:"column:nonce;type:varchar(200)" json:"nonce"`
	ReceiptsRoot      string     `gorm:"column:receipts_root;type:varchar(200)" json:"receipts_root"`
	Sha3Uncles        string     `gorm:"column:sha3_uncles;type:varchar(200)" json:"sha3_uncles"`
	Size              int64      `gorm:"column:size;type:bigint;not null" json:"size"`
	StateRoot         string     `gorm:"column:state_root;type:varchar(200)" json:"state_root"`
	TransactionsRoot  string     `gorm:"column:transactions_root;type:varchar(200)" json:"transactions_root"`
	TransactionsCount int64      `gorm:"column:transactions_count;type:bigint;not null" json:"transactions_count"`
	UnclesCount       int64      `gorm:"column:uncles_count;type:bigint;not null" json:"uncles_count"`
	ChainID           int64      `gorm:"column:chain_id;type:bigint;not null" json:"chain_id"`
	CreatedAt         *time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:Create Time" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:Update Time" json:"updated_at"`
}

// TableName Block's table name
func (*Block) TableName() string {
	return TableNameBlock
}
