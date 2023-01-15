package internal

import (
	"context"
	"log"
	pb "order-service/pkg/v1"
	"sync"

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
