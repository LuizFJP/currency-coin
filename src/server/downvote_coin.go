package main

import (
	"context"
	"fmt"

	pb "LuizFJP/currency-coin-grpc-go/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DownvoteCoin(ctx context.Context, in *pb.CoinRequest) (*pb.CoinResponse, error) {
	result := &CoinItem{}
	data := &CoinItem {
		Name: in.Name,
	}

	err := data.UpdateValidate()
	if err != nil {
		return nil, fmt.Errorf(
			codes.InvalidArgument.String(),
			fmt.Sprint(err),
		)
	}
	err = UpdateByName(data.Name, -1)

	if err != nil {
		return nil, err
	}
	
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: data.Name}}).Decode(&result)
	
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	return &pb.CoinResponse{
		Name: result.Name,
		Price: result.Price,
		Vote: result.Vote,
		
	}, nil
}
