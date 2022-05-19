package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "LuizFJP/currency-coin-grpc-go/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CurrencyCoinServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/coinsdb?retryWrites=true&w=majority"))
	
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Currency Coin Service Started")
	collection = client.Database("coinsdb").Collection("coins")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterCurrencyCoinServiceServer(s, &Server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err = s.Serve(lis); err != nil {
			log.Fatalf("Failed to server %v\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB: %v", err)
	}

	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}