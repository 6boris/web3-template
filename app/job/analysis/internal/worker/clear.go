package worker

import (
	"context"
	"time"
	"web3/app/job/analysis/internal/dao"
)

func (s *Server) _cronClearBlockData() {
	ctx := context.Background()
	_, _ = dao.GlobalDao.MySQLDeleteBlockHistoryData(ctx, time.Now().Add(-time.Hour*48))
	_, _ = dao.GlobalDao.MySQLDeleteBlockTransactionHistoryData(ctx, time.Now().Add(-time.Hour*8))
}
