package defi

import (
	"context"
	"fmt"
	pb "web3/api/interface/web3"
	"web3/app/interface/web3/internal/client"
)

func (s *DefiService) GetTokenPrice(ctx context.Context, req *pb.GetTokenPriceRequest) (*pb.GetTokenPriceReply, error) {
	resp := &pb.GetTokenPriceReply{
		Price: make(map[string]*pb.GetTokenPriceReply_Price, 0),
	}
	sourceToken := req.GetSourceToken()
	targetToken := req.GetTargetToken()

	if sourceToken == "" || targetToken == "" {
		sourceToken = "ETH"
		targetToken = "USD"
	}
	data, err := client.GlobalClient.DefiGetPrice(ctx, sourceToken, targetToken)
	if err != nil {
		return nil, err
	}
	resp.Price[fmt.Sprintf("%s_%s", sourceToken, targetToken)] = &pb.GetTokenPriceReply_Price{
		Price:       data.GetPrice(),
		Timestamp:   data.GetTimestamp(),
		SourceToken: data.GetSourceToken(),
		TargetToken: data.GetTargetToken(),
	}
	return resp, nil
}
