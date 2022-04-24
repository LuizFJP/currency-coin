package main

import (
	"context"
	"fmt"
	"log"

	// "log"

	pb "LuizFJP/currency-coin-grpc-go/proto"
	"github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a CoinItem) CreateValidate() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Price, validation.Required),
	)
}

func (s *Server) CreateCoin(ctx context.Context, in *pb.CreateCoinRequest) (*pb.CoinResponse, error) {
	
	data := &CoinItem {
		Name: in.Name,
		Price: in.Price,
		Vote: 0,
	}
	err := data.CreateValidate()
	log.Print(err)
	// validate:= validator.New()
	// err := validate.Struct(data)
	// validationErrors := err.(validator.ValidationErrors)

	// log.Println(validationErrors.Error())
	// if validationErrors != nil {
	// 	return nil, status.Errorf(
	// 		codes.InvalidArgument,
	// 		fmt.Sprint(validationErrors.Error()),
	// 	)
	// }

	_, err = collection.InsertOne(ctx, data)

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
