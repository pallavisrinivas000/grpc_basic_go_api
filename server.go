package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	crud "/Users/pallavi/projects/grpc_basic_go_api/proto"
)

type server struct {
}

func (s *server) CreateItem(ctx context.Context, req *crud.CreateItemRequest) (*crud.Item, error) {
	newItem := &crud.Item{
		Id:    "1",
		Name:  req.GetName(),
		Price: req.GetPrice(),
	}
	return newItem, nil
}

func (s *server) GetItem(ctx context.Context, req *crud.GetItemRequest) (*crud.Item, error) {
	return &crud.Item{
		Id:    req.Id,
		Name:  "Sample Item",
		Price: 10.99,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	crud.RegisterItemServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
