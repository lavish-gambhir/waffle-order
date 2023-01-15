package internal

import (
	"context"
	"fmt"
	"log"
	pb "order-client/pkg/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type OrderClient struct {
	conn *grpc.ClientConn
	cl   pb.OrderManagementClient
}

func NewOrderClient(addr string) *OrderClient {
	c, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect to grpc server: %v", err)
	}
	return &OrderClient{conn: c, cl: pb.NewOrderManagementClient(c)}
}

func (ordr *OrderClient) GetWaffleOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	order, err := ordr.cl.GetOrder(ctx, orderId)
	if err != nil {
		return nil, fmt.Errorf("unable to get the order for id %v: %v", orderId, err)
	}
	log.Printf("order = %v\n", order)
	return order, nil
}

func (o *OrderClient) CloseConnection() {
	o.conn.Close()
}
