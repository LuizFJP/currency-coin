package main

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	
	pb "LuizFJP/currency-coin-grpc-go/proto"
	
)

func (s *Server) ListCoins(_ *pb.ListCoinRequest, stream pb.CurrencyCoinService_ListCoinsServer) error {

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &CoinItem{}
		err := cursor.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while deconding data from MongoDB: %v", err),
			)
		}
		stream.Send(documentToCurrency(data))
	}

		if err = cursor.Err(); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
	
	}
	return nil
}