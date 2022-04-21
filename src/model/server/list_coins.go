package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/LuizFJP/currency-coin-grpc-go/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListCoins(in *emptypb.Empty, stream pb.CurrencyCoinService_ListCoinsServer) error {
	log.Println("ListCoins was invoked")

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

defer cursor.Close(context.Background())

log.Println(cursor)

	for cursor.Next(context.Background()) {
		data := &CoinItem{}
		err := cursor.Decode(data)
		log.Println(err)
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