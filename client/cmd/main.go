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
	orderId, _ := orderClient.AddWaffleOrder(ctx, sampleWaffleOrders()[1])
	_, err = orderClient.GetWaffleOrder(ctx, &wrapperspb.StringValue{Value: orderId})
	if err != nil {
		log.Fatal(err)
	}
	searchResults := orderClient.SearchWithWaffleName(ctx, &wrapperspb.StringValue{Value: "oco"})
	fmt.Println(searchResults)
	fmt.Println("main done")
}

// TODO: this will be fetched from db
func sampleWaffleOrders() []*v1.Order {
	return []*v1.Order{
		{
			Id: "1",
			Waffles: []*v1.Waffle{
				{
					Name: "Coco",
					Type: "Belgium",
				},
				{
					Name: "Choco",
					Type: "Not Belgium",
				},
			},
			Description: "Some description",
			Price:       123.0,
		},
		{
			Id: "2",
			Waffles: []*v1.Waffle{
				{
					Name: "Moco",
					Type: "Zelgium",
				},
				{
					Name: "Thoco",
					Type: "Not Zelgium",
				},
			},
			Description: "Some description",
			Price:       223.0,
		},
	}
}
