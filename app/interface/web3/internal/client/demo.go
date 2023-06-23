package client

import (
	"context"
	"github.com/google/uuid"
	pb "web3/api/service/demo"
)

func (c *Client) DemoGetUser(ctx context.Context, id int64) (*pb.UserItem, error) {
	user, err := c.ClientAppServiceDemoUserGRPC.GetUser(ctx, &pb.GetUserRequest{Id: id})
	return user.GetItem(), err
}

func (c *Client) DemoCreateUser(ctx context.Context) (int64, error) {
	user, err := c.ClientAppServiceDemoUserGRPC.CreateUser(ctx, &pb.CreateUserRequest{
		Item: &pb.UserItem{
			NickName:  uuid.NewString(),
			AvatarUrl: uuid.NewString(),
			Email:     uuid.NewString(),
		},
	})
	return user.GetId(), err
}
