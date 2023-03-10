package internal

import (
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"io"
	"log"
	pb "order-service/pkg/v1"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type OrderServer struct {
	mu       sync.RWMutex
	orderMap map[string]*pb.Order
	pb.UnimplementedOrderManagementServer
}

func NewOrderServer() *OrderServer { return &OrderServer{orderMap: make(map[string]*pb.Order)} }

// TODO: Add tracing support
func (srv *OrderServer) GetOrder(ctx context.Context, id *wrapperspb.StringValue) (*pb.Order, error) {
	if id.Value == "-1" {
		log.Printf("order id is invalid, id: %v", id.Value)
		errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
		ds, err := errorStatus.WithDetails(
			&errdetails.BadRequest_FieldViolation{Field: "ID", Description: fmt.Sprintf("order id received is not valid %s", id.GetValue())},
		)
		if err != nil {
			return nil, errorStatus.Err()
		}
		return nil, ds.Err()
	}
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	log.Printf("request to get order with id:%v\n", id)
	if ord, ok := srv.orderMap[id.GetValue()]; ok {
		return ord, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Unable to find order for id:%v", id.GetValue())
}

func (srv *OrderServer) AddOrder(ctx context.Context, order *pb.Order) (*wrapperspb.StringValue, error) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	orderId, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("unable to gen. order uuid:%v", err)
	}
	order.Id = orderId.String()
	srv.orderMap[order.Id] = order
	return &wrapperspb.StringValue{Value: order.Id}, nil
}

func (srv *OrderServer) SearchWithWaffleName(searchStr *wrapperspb.StringValue, outStream pb.OrderManagement_SearchWithWaffleNameServer) error {
	log.Printf("got search request with name: %v\n", searchStr.GetValue())
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, ordr := range srv.orderMap {
		for _, waffle := range ordr.Waffles {
			if strings.Contains(waffle.Name, searchStr.GetValue()) {
				err := outStream.Send(ordr)
				if err != nil {
					return fmt.Errorf("unable to send order to stream: %v", err)
				}
			}
		}
	}
	return nil
}

func (srv *OrderServer) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	log.Printf("got update order request")
	var orderIds strings.Builder
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapperspb.StringValue{Value: fmt.Sprintf("orders processed! ids: %s", orderIds.String())})
		}
		srv.orderMap[order.Id] = order
		log.Printf("order updated, with id: %s", order.Id)
		orderIds.WriteString(fmt.Sprintf("%s, ", order.Id))
	}
}

func OrderInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Println("====== [order interceptor] ======", info.FullMethod)
	res, err := handler(ctx, req)
	log.Printf("post proc message:%v\n", res)
	return res, err
}

type WrappedStream struct {
	grpc.ServerStream
}

func NewWrappedStream(ss grpc.ServerStream) *WrappedStream { return &WrappedStream{ss} }

func (w *WrappedStream) RecvMsg(m any) error {
	log.Printf("======= [streaming wrapper] ======\treceived msg (type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *WrappedStream) SendMsg(m any) error {
	log.Printf("====== [streaming wrapper] ======\tsend msg (type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func OrderStreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("====== [order stream interceptor] =====", info.FullMethod)
	err := handler(srv, NewWrappedStream(ss))
	if err != nil {
		log.Fatal(err)
	}
	return err
}
