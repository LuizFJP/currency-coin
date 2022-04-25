package main

import (
	"context"
	"fmt"
	"log"

	// "log"

	pb "LuizFJP/currency-coin-grpc-go/proto"

	"github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a CoinItem) CreateValidate() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Price, validation.Required),
	)
}

func checkCoin(name string) error {
	result := &CoinItem{}
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	collection.FindOne(context.TODO(), filter).Decode(&result)
	
	log.Print("test", result)
	if result.Name == name {
		return status.Error(
			codes.AlreadyExists,
			"Can't create coin with an existing name",
		)
	}

	return nil
}

func (s *Server) CreateCoin(ctx context.Context, in *pb.CreateCoinRequest) (*pb.CoinResponse, error) {
	
	data := &CoinItem {
		Name: in.Name,
		Price: in.Price,
		Vote: 0,
	}
	err := data.CreateValidate()
	log.Print(err)

	err = checkCoin(data.Name)

	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
log.Print(data)
	return &pb.CoinResponse{
		Name: data.Name,
		Price: data.Price,
		Vote: data.Vote,
	}, nil
}
