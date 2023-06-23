package defi

import (
	"context"
	pb "web3/api/service/data"
	"web3/app/service/data/internal/dao"
)

func (s *DefiService) GetDefiPrice(ctx context.Context, req *pb.GetDefiPriceRequest) (*pb.GetDefiPriceReply, error) {
	resp := &pb.GetDefiPriceReply{}
	price, err := dao.GlobalDao.BizDefiGetPrice(ctx, req.GetSourceToken(), req.GetTargetToken())
	if err != nil {
		return nil, err
	}
	resp.Item = &pb.DefiPriceItem{
		Price:       price.Price,
		Timestamp:   price.Timestamp,
		SourceToken: price.SourceToken,
		TargetToken: price.TargetToken,
	}
	return resp, nil
}
