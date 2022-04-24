package main

import (
	pb "LuizFJP/currency-coin-grpc-go/proto"
)

type CoinItem struct {
	Name 		string 		`bson:"name" json:"name"`
	Price 	float64 	`bson:"price" json:"price"`
	Vote 		int64 		`bson:"vote" json:"vote"`
}

func documentToCurrency(data *CoinItem) *pb.CoinResponse {
	return &pb.CoinResponse{
		Name: data.Name,
		Price: data.Price,
		Vote: data.Vote,
	}
}