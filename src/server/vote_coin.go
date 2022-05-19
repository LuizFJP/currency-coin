package main

import (
	pb "LuizFJP/currency-coin-grpc-go/proto"
	"context"
	// "fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "google.golang.org/grpc/codes"
)

func (s *Server) VoteCoin(in *pb.CoinRequest, stream pb.CurrencyCoinService_VoteCoinServer) error {
	
	result := &CoinItem{}

	ticker := time.NewTicker(5000 * time.Millisecond)
	done := make(chan bool)

		for {
				select {
				case <-done:
						return nil
				case <-ticker.C:
					err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: in.Name}}).Decode(&result)

					if err != nil {
						return nil
					}

						stream.Send(documentToCurrency(result))

					}
		}

	}


