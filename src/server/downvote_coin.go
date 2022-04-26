package main

import (
	"context"
	"fmt"

	pb "LuizFJP/currency-coin-grpc-go/proto"

	"google.golang.org/grpc/codes"

)

func (s *Server) DownvoteCoin(ctx context.Context, in *pb.CoinRequest) (*pb.CoinResponse, error) {
	result, err := Vote(in, -1)
	if err != nil {
		return nil, fmt.Errorf(
			codes.InvalidArgument.String(),
			fmt.Sprint(err),
		)
	}

	return &pb.CoinResponse{
		Name: result.Name,
		Price: result.Price,
		Vote: result.Vote,
		
	}, nil
}
