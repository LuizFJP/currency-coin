package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/LuizFJP/currency-coin-grpc-go/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoin(ctx context.Context, in *pb.CreateCoinRequest) (*pb.CoinResponse, error) {
	log.Println("Create was invoked")
	data := CoinItem {
		Name: in.Name,
		Price: in.Price,
		Vote: 0,
	}

	_, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &pb.CoinResponse{
		Name: data.Name,
		Price: data.Price,
		Vote: data.Vote,
	}, nil
}
