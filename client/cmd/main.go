package main

import (
	"context"
	"fmt"
	"log"
	"order-client/internal"
	v1 "order-client/pkg/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const addr = "localhost:9090"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect to grpc server: %v", err)
	}
	defer conn.Close()
	orderClient := internal.NewOrderClient(addr, conn)
	helloClient := v1.NewSayHelloClient(conn) // as this is just for test, not creating a wrapper over it.

	resp, err := helloClient.Hello(ctx, &v1.HelloRequest{Who: "order-client"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.String())

	_, err = orderClient.GetWaffleOrder(ctx, &wrapperspb.StringValue{Value: "<id>"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main done")
}
