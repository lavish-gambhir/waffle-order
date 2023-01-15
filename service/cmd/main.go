package main

import (
	"fmt"
	"log"
	"net"
	v1 "order-service/pkg/v1"

	"order-service/internal"

	"google.golang.org/grpc"
)

const port = "9090"

func main() {
	lsnr, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("unable to listen on port %q: %v", port, err)
	}
	srv := grpc.NewServer()
	v1.RegisterOrderManagementServer(srv, internal.NewOrderServer())
	fmt.Printf("server listening on port:%s\n", port)
	if err := srv.Serve(lsnr); err != nil {
		log.Fatalf("unable to serve %v", err)
	}
}
