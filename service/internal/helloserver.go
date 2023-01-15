package internal

import (
	"context"
	"fmt"
	v1 "order-service/pkg/v1"
)

type HelloServer struct {
	v1.UnimplementedSayHelloServer
}

func NewHelloServer() *HelloServer { return &HelloServer{} }

func (h *HelloServer) Hello(ctx context.Context, r *v1.HelloRequest) (*v1.HelloResponse, error) {
	resp := fmt.Sprintf("Dear %s, request received successfully :)!", r.Who)
	return &v1.HelloResponse{Response: resp}, nil
}
