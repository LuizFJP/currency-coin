package main

import (
	pb "LuizFJP/currency-coin-grpc-go/proto"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
)

func (s *Server) Delete(ctx context.Context, in *pb.CoinRequest) (*pb.DeleteResponse, error) {

	data := &CoinItem{
		Name: in.Name,
	}
	
	err := data.NameValidate()
	if err != nil {
		return nil, fmt.Errorf(
			codes.InvalidArgument.String(),
			fmt.Sprint(err),
		)
	}

	filter := bson.D{primitive.E{Key: "name", Value: data.Name}}
	
	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf(
			codes.InvalidArgument.String(),
			fmt.Sprint(err),
		)
	}

	if result.DeletedCount == 0 {
		return nil, fmt.Errorf(
			codes.NotFound.String(),
			fmt.Sprintf("%v was not found, can't be possible to delete", data.Name),
		)
	}

	msg := data.Name + " was deleted successful!"

	return &pb.DeleteResponse{
		Message: msg,
	}, nil

}