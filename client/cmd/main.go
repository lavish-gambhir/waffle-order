package main

import (
	"context"
	"fmt"
	"log"
	"order-client/internal"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

const addr = "localhost:9090"

func main() {
	orderClient := internal.NewOrderClient(addr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer orderClient.CloseConnection()

	_, err := orderClient.GetWaffleOrder(ctx, &wrapperspb.StringValue{Value: "<id>"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main done")
}
