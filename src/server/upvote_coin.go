package main

import (
	"context"
	"fmt"

	pb "LuizFJP/currency-coin-grpc-go/proto"

	"github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a CoinItem) UpdateValidate() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
	)
}

func UpdateByName(name string) (error) {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "vote", Value: 1}}}}
	
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return status.Error(
			codes.Internal,
			"It's not possible to upvote. Coin doesn't exist",
		)
	}
	return nil

}

func (s *Server) UpvoteCoin(ctx context.Context, in *pb.CoinRequest) (*pb.CoinResponse, error) {
	result := &CoinItem{}
	data := CoinItem {
		Name: in.Name,
	}

	err := data.UpdateValidate()
	if err != nil {
		return nil, fmt.Errorf(
			codes.InvalidArgument.String(),
			fmt.Sprint(err),
		)
	}
	err = UpdateByName(data.Name)

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
