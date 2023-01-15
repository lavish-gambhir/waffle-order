package internal

import (
	"context"
	"fmt"
	"log"
	pb "order-client/pkg/v1"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type OrderClient struct {
	cl pb.OrderManagementClient
}

func NewOrderClient(addr string, conn *grpc.ClientConn) *OrderClient {
	return &OrderClient{cl: pb.NewOrderManagementClient(conn)}
}

func (ordr *OrderClient) GetWaffleOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	order, err := ordr.cl.GetOrder(ctx, orderId)
	if err != nil {
		return nil, fmt.Errorf("unable to get the order for id %v: %v", orderId, err)
	}
	log.Printf("order = %v\n", order)
	return order, nil
}
