package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "order-client/pkg/v1"
	"time"

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

func (ordr *OrderClient) AddWaffleOrder(ctx context.Context, order *pb.Order) (string, error) {
	res, err := ordr.cl.AddOrder(ctx, order)
	if err != nil {
		log.Fatalf("unable to add a new order: %v\n", err)
	}
	log.Printf("added order, id: %v\n", res.GetValue())
	return res.GetValue(), nil
}

func (ordr *OrderClient) SearchWithWaffleName(ctx context.Context, query *wrapperspb.StringValue) []*pb.Order {
	var results []*pb.Order
	searchStream, _ := ordr.cl.SearchWithWaffleName(ctx, query)
	for {
		searchResult, err := searchStream.Recv()
		if err == io.EOF {
			// got all the search results.. we are done!
			break
		}
		if err != nil {
			log.Printf("error while reaciving results of search query: %v\n", err)
			break
		}
		log.Print("Search result: ", searchResult)
		results = append(results, searchResult)
	}
	return results
}

func (ordr *OrderClient) UpdateOrders(orders []*pb.Order, ctx context.Context) (string, error) {
	outStream, err := ordr.cl.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("unable to update orders:%v", err)
	}
	if err = outStream.Send(orders[0]); err != nil {
		log.Fatalf("unable to send order:%v", err)
	}
	if err = outStream.Send(orders[1]); err != nil {
		log.Fatalf("unable to send order:%v", err)
	}
	updateResult, err := outStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("closeAndReceive error:%v", err)
	}
	return updateResult.GetValue(), nil
}

func OrderUnaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Println("method:", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Println(reply)
	return err
}

type WrappedStream struct {
	grpc.ClientStream
}

func NewOrderStreamingInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	log.Println("===== [Client interceptor] =====", method)
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

func (w *WrappedStream) RecvMsg(m any) error {
	log.Printf("====== [client::order stream interceptor\trecieved msg (type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *WrappedStream) SendMsg(m any) error {
	log.Printf("====== [client::order stream interceptor\trecieved msg (type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func newWrappedStream(c grpc.ClientStream) *WrappedStream { return &WrappedStream{c} }
