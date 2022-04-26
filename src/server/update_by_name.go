package main

import (
	"context"

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