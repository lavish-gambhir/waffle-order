package internal

import (
	pb "order-service/pkg/v1"
)

type OrderServer struct {
	orderMap map[string]*pb.Order
}

func NewOrderServer() *OrderServer { return &OrderServer{orderMap: make(map[string]*pb.Order)} }
