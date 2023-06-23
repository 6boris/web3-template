package client

import (
	"context"
	pb "web3/api/service/data"
)

func (c *Client) DefiGetPrice(ctx context.Context, sourceToken, targetToken string) (*pb.DefiPriceItem, error) {
	data, err := c.ClientAppServiceDataDefiGRPC.GetDefiPrice(ctx, &pb.GetDefiPriceRequest{SourceToken: sourceToken, TargetToken: targetToken})
	return data.GetItem(), err
}
