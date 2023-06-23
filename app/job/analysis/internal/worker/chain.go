package worker

import (
	"context"
	"web3/app/job/analysis/internal/dao"
)

func (s *Server) _cronSyncChainInfo() {
	chainInfo, err := dao.GlobalDao.MySQLGetAllChainInfoList(context.Background())
	if err != nil {
		return
	}
	for _, chain := range chainInfo {
		s._localCacheChainInfo.Store(chain.UniqueID, chain)
	}
}
