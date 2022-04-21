package main

import (
	pb "github.com/LuizFJP/currency-coin-grpc-go/proto"
)

type CoinItem struct {
	Name string `bson:"name"`
	Price float64 `bson: "price"`
	Vote int64 `bson: "vote"`
}

func documentToCurrency(data *CoinItem) *pb.CoinResponse {
	return &pb.CoinResponse{
		Name: data.Name,
		Price: data.Price,
		Vote: data.Vote,
	}
}