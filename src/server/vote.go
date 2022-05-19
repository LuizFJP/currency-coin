package main

import (
	"context"
	"fmt"

	pb "LuizFJP/currency-coin-grpc-go/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UpdateByName(name string, increment int) (error) {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "vote", Value: increment}}}}
	
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return status.Error(
			codes.Internal,
			"It's not possible to upvote. Coin doesn't exist",
		)
	}
	return nil
}

func Vote(in *pb.CoinRequest, increment int) (*CoinItem, error){
	result := &CoinItem{}
	data := &CoinItem {
		Name: in.Name,
	}

	err := data.NameValidate()
	if err != nil {
		return nil, fmt.Errorf(
			codes.InvalidArgument.String(),
			fmt.Sprint(err),
		)
	}
	err = UpdateByName(data.Name, increment)

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

	return result, nil
}