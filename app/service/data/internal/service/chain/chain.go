package chain

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-module/carbon/v2"
	pb "web3/api/service/data"
	"web3/app/service/data/internal/dao"
	"web3/app/service/data/storage/mysql/model"
)

func (s ChainService) SyncBlockInfo(ctx context.Context, req *pb.SyncBlockInfoRequest) (*pb.SyncBlockInfoReply, error) {
	resp := &pb.SyncBlockInfoReply{}
	blockTime := carbon.CreateFromTimestamp(req.GetItem().GetTimestamp()).ToStdTime()
	err := dao.GlobalDao.MySQLSyncBlock(ctx, &model.Block{
		Number:            req.GetItem().GetNumber(),
		Hash:              req.GetItem().GetHash(),
		ParentHash:        req.GetItem().GetParentHash(),
		Timestamp:         &blockTime,
		Difficulty:        req.GetItem().GetDifficulty(),
		ExtraData:         req.GetItem().GetExtraData(),
		GasLimit:          req.GetItem().GetGasLimit(),
		GasUsed:           req.GetItem().GetGasUsed(),
		BaseFeePerGas:     req.GetItem().GetBaseFeePerGas(),
		Miner:             req.GetItem().GetBaseFeePerGas(),
		MixHash:           req.GetItem().GetMixHash(),
		Nonce:             req.GetItem().GetNonce(),
		ReceiptsRoot:      req.GetItem().GetReceiptsRoot(),
		Sha3Uncles:        req.GetItem().GetSha3Uncles(),
		Size:              req.GetItem().GetSize(),
		StateRoot:         req.GetItem().GetStateRoot(),
		TransactionsRoot:  req.GetItem().GetTransactionsRoot(),
		TransactionsCount: req.GetItem().GetTransactionsCount(),
		UnclesCount:       req.GetItem().GetUnclesCount(),
		ChainID:           req.GetItem().GetChainId(),
		ChainUniqueID:     req.GetItem().GetChainUniqueId(),
	})
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	return resp, nil
}

func (s ChainService) SyncBlockTransactionInfo(ctx context.Context, req *pb.SyncBlockTransactionInfoRequest) (*pb.SyncBlockTransactionInfoReply, error) {
	blockTime := carbon.CreateFromTimestamp(req.GetItem().GetBlockTimestamp()).ToStdTime()
	resp := &pb.SyncBlockTransactionInfoReply{}
	err := dao.GlobalDao.MySQLSyncBlockTransaction(ctx, &model.BlockTransaction{
		Type:                 req.GetItem().GetType(),
		Status:               req.GetItem().GetStatus(),
		BlockHash:            req.GetItem().GetBlockHash(),
		BlockNumber:          req.GetItem().GetBlockNumber(),
		BlockTimestamp:       &blockTime,
		TransactionHash:      req.GetItem().GetTransactionHash(),
		TransactionIndex:     req.GetItem().GetTransactionIndex(),
		FromAddress:          req.GetItem().GetFromAddress(),
		ToAddress:            req.GetItem().GetToAddress(),
		Value:                req.GetItem().GetValue(),
		Input:                req.GetItem().GetInput(),
		Nonce:                req.GetItem().GetNonce(),
		ContractAddress:      req.GetItem().GetContractAddress(),
		Gas:                  req.GetItem().GetGas(),
		GasPrice:             req.GetItem().GetGasPrice(),
		GasUsed:              req.GetItem().GetGasUsed(),
		EffectiveGasPrice:    req.GetItem().GetEffectiveGasPrice(),
		CumulativeGasUsed:    req.GetItem().GetCumulativeGasUsed(),
		MaxPriorityFeePerGas: req.GetItem().GetMaxPriorityFeePerGas(),
		MaxFeePerGas:         req.GetItem().GetMaxFeePerGas(),
		R:                    req.GetItem().GetR(),
		S:                    req.GetItem().GetS(),
		V:                    req.GetItem().GetV(),
		ChainID:              req.GetItem().GetChainId(),
		ChainUniqueID:        req.GetItem().GetChainUniqueId(),
	})
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	return resp, nil
}
