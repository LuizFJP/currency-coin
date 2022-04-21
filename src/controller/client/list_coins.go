package main

import (
	"context"
	"io"
	"log"

	pb "github.com/LuizFJP/currency-coin-grpc-go/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listCoins(c pb.CurrencyCoinServiceClient) {
	log.Println(("listCoins was invoked"))

	stream, err := c.ListCoins(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling listCoins: %v\n", err)
	}
	
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something Happened: %v\n", err)
		}

		log.Println(res)
	}
}