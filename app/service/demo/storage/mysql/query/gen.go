// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:               db,
		Block:            newBlock(db, opts...),
		BlockTransaction: newBlockTransaction(db, opts...),
		ChainClient:      newChainClient(db, opts...),
		ChainInfo:        newChainInfo(db, opts...),
		ChainMetric:      newChainMetric(db, opts...),
		User:             newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Block            block
	BlockTransaction blockTransaction
	ChainClient      chainClient
	ChainInfo        chainInfo
	ChainMetric      chainMetric
	User             user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		Block:            q.Block.clone(db),
		BlockTransaction: q.BlockTransaction.clone(db),
		ChainClient:      q.ChainClient.clone(db),
		ChainInfo:        q.ChainInfo.clone(db),
		ChainMetric:      q.ChainMetric.clone(db),
		User:             q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		Block:            q.Block.replaceDB(db),
		BlockTransaction: q.BlockTransaction.replaceDB(db),
		ChainClient:      q.ChainClient.replaceDB(db),
		ChainInfo:        q.ChainInfo.replaceDB(db),
		ChainMetric:      q.ChainMetric.replaceDB(db),
		User:             q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Block            *blockDo
	BlockTransaction *blockTransactionDo
	ChainClient      *chainClientDo
	ChainInfo        *chainInfoDo
	ChainMetric      *chainMetricDo
	User             *userDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Block:            q.Block.WithContext(ctx),
		BlockTransaction: q.BlockTransaction.WithContext(ctx),
		ChainClient:      q.ChainClient.WithContext(ctx),
		ChainInfo:        q.ChainInfo.WithContext(ctx),
		ChainMetric:      q.ChainMetric.WithContext(ctx),
		User:             q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
