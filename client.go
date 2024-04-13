package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	crud "grpc_basic_go_api/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := crud.NewItemServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Create an item
	item, err := client.CreateItem(context.Background(), &crud.CreateItemRequest{
		Name:  "Sample Item",
		Price: 10.99,
	})
	if err != nil {
		log.Fatalf("could not create item: %v", err)
	}
	log.Printf("Created item: %v", item)

	// Get an item
	getItem, err := client.GetItem(context.Background(), &crud.GetItemRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not get item: %v", err)
	}
	log.Printf("Got item: %v", getItem)
}
