package main

import (
	pb "LuizFJP/currency-coin-grpc-go/proto"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDelete(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	client := pb.NewCurrencyCoinServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success to delete coin", func (mt *mtest.T)  {
		collection = mt.Coll
		fakeCoin := &CoinItem {
			Name: "BTC",
		}

		mt.AddMockResponses(bson.D{
			primitive.E{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1},
		})

		req := &pb.CoinRequest{Name: fakeCoin.Name}
		_, err := client.Delete(context.Background(), req)

		if err != nil {
			t.Errorf("Something went wrong: %v", err)
		}
	})
}

func TestFailed(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	client := pb.NewCurrencyCoinServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("It's not possible to delete a coin with name not existed", func (mt *mtest.T)  {
		collection = mt.Coll
		fakeCoin := &CoinItem {
			Name: "BTC",
		}

		mt.AddMockResponses(bson.D{
			primitive.E{Key: "ok", Value: 1}, {Key: "acknowledged", Value: false}, {Key: "n", Value: 0},
		})

		req := &pb.CoinRequest{Name: fakeCoin.Name}
		_, err := client.Delete(context.Background(), req)

		if err == nil {
			t.Errorf("Something went wrong: %v", err)
		}
	})
}