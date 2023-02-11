package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "order-service/pkg/v1"
	"strings"
	"sync"

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
